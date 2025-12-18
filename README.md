
# SLG 服务器

## 开发环境
* golang 1.25
* mac mini

## 目录说明
| 目录                 | 说明           |
|--------------------|--------------|
| cmd                | 应用程序入口       |
| data               | 地图数据, 行军测试数据 |
| internal           | 内部服务逻辑       |
| internal.bootstrap | 启动引导         |
| internal.battle    | 战斗           |
| internal.config    | 配置数据结构       |
| internal.march     | 行军           |
| internal.node      | 节点信息         |
| internal.player    | 玩家           |
| internal.world     | 世界           |
| opt                | 启动脚本         |
| pkg                | 可复用工具入口      |

## 启动服务器
需要安装**go**环境
```bash
./opt/start.sh
```


## TODO
- [x] 定时器: 1 tick/s
- [x] World: 整个地图管理
- [x] Graph: 最基础的无权无向图
- [x] MarchMgr: 行军
- [ ] PlayerMgr: 玩家管理
- [ ] BattleUnit: 战斗单位
