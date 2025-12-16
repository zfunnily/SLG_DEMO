
# SLG 服务器

## 开发环境
* golang 1.25
* mac mini

## 目录说明
| 目录                 | 说明      |
|--------------------|---------|
| cmd                | 应用程序入口  |
| data               | 地图数据    |
| internal           | 内部服务逻辑  |
| internal.battle    | 战斗      |
| internal.bootstrap | 启动引导    |
| internal.config    | 配置数据结构  |
| internal.march     | 行军管理    |
| internal.node      | 节点信息    |
| internal.player    | 玩家      |
| internal.world     | 世界      |
| opt                | 启动脚本    |
| pkg                | 可复用工具入口 |

## TODO
- [x] 定时器: 1 tick/s
- [ ] SLG_MAP: 整个地图管理
- [ ] Graph: 最基础的无权无向图
- [ ] NodeInfo: 节点信息管理
- [ ] MarchMgr: 行军
- [ ] PlayerMgr: 玩家管理
- [ ] BattleUnit: 战斗单位
