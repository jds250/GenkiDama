import math

import matplotlib.pyplot as plt

#
# 解析链上
#

def add_on_chain_cost(op, onChainCost, blockNum):
    global OnCostMap
    global OnStartBlock
    global OnEndBlock
    global ArguCnt

    if op in OnCostMap:
        OnCostMap[op] = (float(OnCostMap[op]) + float(onChainCost)) / 2
    else:
        OnCostMap[op] = float(onChainCost)

    if op == 'ChallengeCreate':
        OnEndBlock = blockNum
    if op == 'AddCheckPoint':
        OnStartBlock = blockNum
    if op == 'QuestionArgu':
        ArguCnt += 1


def add_off_chain_cost(op, offChainCost):
    global OffCostMap

    if op in OffCostMap:
        OffCostMap[op] = (float(OffCostMap[op]) + float(offChainCost)) / 2
    else:
        OffCostMap[op] = offChainCost


# 创建数据容器
OffCostMap = {}
OnCostMap = {}
OnStartBlock = 0
OnEndBlock = 0
ArguCnt = 0

# 读取数据
with  open('TxLogInfo.txt') as f:
    lines = f.readlines()
    # print(lines)
    f.close()

print("lines len = ", len(lines))

# 解析数据
for line in lines:
    elems = line.split(" ")
    op = elems[4]
    onChainCost = elems[5]
    blockNum = elems[6][:len(elems[6]) - 1]
    add_on_chain_cost(op, onChainCost, int(blockNum))

# 输出解析结果
print("ON CHAIN COST:", end="(")
print("Start:", OnStartBlock, "End:", OnEndBlock, ")")
for key in OnCostMap:
    print(" - ", key, OnCostMap[key], "gas")

#
# 解析链下
#
# 读取数据
with  open('OffChainInfo.txt') as f:
    lines = f.readlines()
    # print(lines)
    f.close()


# 解析数据
for line in lines:
    elems = line.split(" ")
    op = elems[4]
    timeCost = elems[5]
    if timeCost[len(timeCost)-3] == 'm':
        timeCost = float(timeCost[:len(timeCost)-3])
    else:
        timeCost = float(timeCost[:len(timeCost)-2])*1000
    # print("op=", op, "\ttimeCost=", timeCost)
    add_off_chain_cost(op, timeCost)

# 输出解析结果
print("OFF CHAIN COST:")
for key in OffCostMap:
    print(" - ", key, OffCostMap[key], "ms")


#
# 细粒度分析：交互式挑战
#

# 链下开销：两次证明生成的开销
proofGenerate = OffCostMap['ChallengeCreate'] + OffCostMap['ChallengeFinalProof']
argu_offchain = (OffCostMap['QuestionArgu'] + OffCostMap['ChallengeResponse'])
all_off_chain_cost = proofGenerate + argu_offchain*ArguCnt
print("链下开销汇总 = ", proofGenerate, "+", argu_offchain*ArguCnt, " = ", all_off_chain_cost)

# 验证证明的开销
print("### 交互式bridge")
proofVerify = OnCostMap['ChallengeCreate'] + OnCostMap['ChallengeFinalProof']
# 交互式挑战的开销
arguCost = OnCostMap['QuestionArgu']
respCost = OnCostMap['ChallengeResponse']
argu_onchain = (arguCost + respCost)
all_on_chain_cost = argu_onchain*ArguCnt + proofVerify

print("链上开销汇总 = ", proofVerify, "+", argu_onchain*ArguCnt, " = ", all_on_chain_cost)

#
# 细粒度分析：传统挑战
#
block_cnt = OnEndBlock - OnStartBlock
old_offchain_cost = block_cnt * OffCostMap['AddCheckPoint']
old_onchain_cost = block_cnt * OnCostMap['AddCheckPoint']
print("### 传统bridge")
print(" ` block count =", block_cnt)
print(" · 传统链下开销 =", old_offchain_cost)
print(" · 传统链上开销 =", old_onchain_cost)


#
# 绘制图像
# 图1-2：各类操作的开销：链下生成证明，链上证明验证，一轮交互式验证开销，总开销
#
op_on_chain_label = ['Proof Verify', 'One Argu', 'All On Chain']
op_on_chain_cost = [proofVerify, argu_onchain, all_on_chain_cost]

op_off_chain_label = ['Proof Generate', 'One Argu', 'All Off Chain']
op_off_chain_cost = [proofGenerate, argu_offchain, all_off_chain_cost]

print(op_on_chain_cost)
plt.bar(op_on_chain_label, op_on_chain_cost)
plt.xlabel('Op Name')
plt.ylabel('On Chain Cost (gas)')
plt.savefig('./op_on_chain.jpg')
plt.show()

print(op_off_chain_cost)
plt.bar(op_off_chain_label, op_off_chain_cost)
plt.xlabel('Op Name')
plt.ylabel('Off Chain Cost (ms)')
plt.savefig('./op_off_chain.jpg')
plt.show()

#
# 绘制图像
# 图3：随着区块的增加，验证成本的变换
#
block_x = [10, 20, 30, 40, 50]
block_on_chain_cost = []
tradition_on_chain_cost = []
block_off_chain_cost = []
tradition_off_chain_cost = []
for i in block_x:
    n = (int(math.log2(i))+1)
    on_cost = proofVerify + argu_onchain * n
    off_cost = (proofGenerate + argu_offchain * n)/1000

    tradition_on = proofVerify*i
    tradition_off = proofGenerate*i/1000

    block_off_chain_cost.append(off_cost)
    block_on_chain_cost.append(on_cost)

    tradition_off_chain_cost.append(tradition_off)
    tradition_on_chain_cost.append(tradition_on)

plt.plot(block_x, block_on_chain_cost, 'b*--', alpha=0.5, linewidth=1, label='On Chain Cost (gas)')
plt.plot(block_x, tradition_on_chain_cost, 'b*--', alpha=0.5, linewidth=1, label='On Chain Cost (gas)')

plt.xlabel('blockNum')
plt.ylabel('On Chain Cost (gas)')
plt.savefig('./num_cost_on_chain2.jpg')
plt.show()

plt.plot(block_x, block_off_chain_cost, 'rs--', alpha=0.5, linewidth=1, label='Off Chain Cost (s)')
plt.plot(block_x, tradition_off_chain_cost, 'rs--', alpha=0.5, linewidth=1, label='Off Chain Cost (s)')
plt.xlabel('blockNum')
plt.ylabel('Off Chain Cost (s)')
plt.savefig('./num_cost_off_chain2.jpg')
plt.show()