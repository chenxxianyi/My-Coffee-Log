# My Coffee Log 产品优化开发任务拆解

> 文档版本：V1.0  
> 文档日期：2026-08-18  
> 方案来源：[PRODUCT_OPTIMIZATION_ROADMAP.md](./PRODUCT_OPTIMIZATION_ROADMAP.md)  
> 适用团队：产品、设计、前端、后端、测试  
> 目标：将产品优化路线图拆解为可估时、可开发、可测试、可验收的任务清单。

---

## 1. 文档使用说明

### 1.1 任务状态

| 状态 | 含义 |
|---|---|
| `TODO` | 尚未开始 |
| `PARTIAL` | 已有部分能力，需要补齐 |
| `DOING` | 开发中 |
| `BLOCKED` | 被依赖或外部条件阻塞 |
| `DONE` | 已开发、测试并验收 |

### 1.2 工作量标记

工作量以单人有效开发日估算，不包含产品评审等待时间：

| 标记 | 参考工作量 |
|---|---:|
| XS | 0.5 天以内 |
| S | 1—2 天 |
| M | 3—5 天 |
| L | 6—10 天 |
| XL | 需要进一步拆分 |

### 1.3 角色缩写

- `PM`：产品经理。
- `UX`：交互/视觉设计。
- `FE`：前端开发。
- `BE`：后端开发。
- `QA`：测试。
- `DA`：数据分析；没有专职角色时由 PM/BE 共同负责。

### 1.4 完成定义 Definition of Done

每个开发任务只有同时满足以下条件才可以标记为 `DONE`：

- 需求与边界已评审。
- 前后端类型和接口契约一致。
- 关键逻辑具有自动化测试或可重复测试步骤。
- 移动端常见尺寸完成视觉检查。
- 错误、空状态、加载和重试状态完整。
- 新增行为已接入埋点或明确说明不需要埋点。
- 不产生新的虚假默认数据。
- TypeScript、Go 测试和生产构建通过。
- 验收标准逐项确认。

---

## 2. 当前代码基线

### 2.1 当前技术结构

- 前端：Vue 3、TypeScript、Pinia、Vue Router、Tailwind CSS、Vite。
- 后端：Go、Gin、GORM、MySQL、Redis。
- 鉴权：JWT。
- AI：服务端 AI Service，支持 Mock 和外部模型。
- 数据迁移：启动时通过 GORM `AutoMigrate` 执行。

### 2.2 已有主要模块

| 模块 | 前端入口 | 后端入口 |
|---|---|---|
| 快捷/精细记录 | `frontend/src/pages/CreateCoffeeLog.vue` | `backend/internal/service/coffee_log_service.go` |
| 记录 API | `frontend/src/api/coffeeLog.ts` | `backend/internal/handler/coffee_log_handler.go` |
| 状态管理 | `frontend/src/stores/coffeeLog.ts` | — |
| 咖迹统计 | `frontend/src/pages/Stats.vue` | `backend/internal/service/stats_service.go` |
| 月度回顾 | `frontend/src/pages/MonthlyReview.vue` | `/api/v1/stats/monthly-review` |
| 门店 | `frontend/src/pages/CoffeeShops.vue` | `/api/v1/coffee-shops` |
| 咖啡豆 | 记录页内选择 | `/api/v1/coffee-beans` |
| AI | 多个页面 | `/api/v1/ai/*` |
| 分享卡 | `frontend/src/pages/ShareCard.vue` | `/api/v1/ai/share-copy` |

### 2.3 已确认的技术问题

1. `coffee_logs` 的六维感官值是非空整数，默认值为 `0`，无法区分“未填写”和“真实评分为 0”。
2. `CreateCoffeeLogRequest` 仍要求 `coffee_name`，快捷模式需要自动生成名称。
3. 快捷模式仍写入：
   - `Local Coffee Spot`；
   - “一杯温润安静的手账记录。”；
   - 默认六维感官分数。
4. 服务端会始终生成 Mock AI 摘要，即使用户没有要求生成 AI 内容。
5. 门店服务会根据 `ShopName` 自动创建/更新门店，虚假默认门店会继续污染门店数据。
6. 统计接口未返回有效样本数和统计时间范围。
7. 咖啡人格规则可以在样本量不足时产生结论。
8. 前端路由全部同步加载，生产包体目前存在大于 500 kB 的警告。
9. 当前缺少统一的产品事件采集层。

---

## 3. 推荐发布节奏

以 2 周一个 Sprint、90 天约 6 个 Sprint 规划：

| Sprint | 周期 | 主题 | 必须交付 |
|---|---|---|---|
| Sprint 0 | 第 1 周 | 基线与数据契约 | 指标定义、埋点壳、数据库方案、测试基线 |
| Sprint 1 | 第 2—3 周 | 数据可信 | 清除虚假默认数据、迁移旧数据、统计样本规则 |
| Sprint 2 | 第 4—5 周 | 首次激活 | 首次引导、第一杯流程、成功反馈 |
| Sprint 3 | 第 6—7 周 | 留存闭环 | 周度回顾、历史记忆、再次冲煮 |
| Sprint 4 | 第 8—9 周 | 复用与 AI 信任 | 门店/豆子复用、AI 数据说明和反馈 |
| Sprint 5 | 第 10—12 周 | 增长与质量 | 分享回流、性能、无障碍、全量验收 |

原则：Sprint 1 完成前，不进入依赖可信统计的洞察开发。

---

## 4. Sprint 0：工程与指标基础

### TASK-000：建立优化版本基线

- 状态：`TODO`
- 优先级：P0
- 角色：PM、FE、BE、QA
- 工作量：S
- 依赖：无

#### 开发步骤

1. 记录当前接口、数据库结构和关键页面截图。
2. 确认开发、测试和生产环境的配置差异。
3. 建立本轮优化功能开关：
   - `data_quality_v2`；
   - `onboarding_v1`；
   - `record_success_v1`；
   - `weekly_review_v1`。
4. 确认灰度策略：内部账号 → 10% 用户 → 50% → 全量。
5. 为每个 Sprint 建立回滚说明。

#### 建议文件

- 新增：`backend/internal/config/feature_flags.go`。
- 新增：`frontend/src/constants/featureFlags.ts`。
- 更新：`backend/internal/config/config.go`。

#### 验收标准

- 功能可以按环境或用户范围开关。
- 关闭开关后仍可使用旧流程。
- 发布说明包含回滚条件和负责人。

### TASK-001：定义核心指标口径

- 状态：`DONE` ✅
- 优先级：P0
- 角色：PM、DA、BE
- 工作量：S
- 依赖：无
- 完成日期：2026-08-18

#### 开发步骤

1. 定义“有效记录”：
   - 真实用户创建；
   - 具备有效日期；
   - 咖啡类型由用户确认；
   - 非测试数据；
   - 不含已知默认污染记录。
2. 定义“有效感官记录”：六个维度均由用户主动提交。
3. 定义首次激活、周活跃记录用户、再次冲煮等指标。
4. 编写指标字典，包括公式、时间窗口、去重规则和异常处理。

#### 交付物

- 新增：`docs/METRICS_DICTIONARY.md`，若暂不新增 `docs` 目录，可放在根目录。

#### 验收标准

- 前端埋点、后端聚合和产品报表使用同一口径。
- 每个指标有唯一名称、负责人和数据来源。

### TASK-002：建立产品事件采集层

- 状态：`TODO`
- 优先级：P0
- 角色：FE、BE、DA
- 工作量：M
- 依赖：TASK-001

#### 推荐方案

先实现与供应商无关的前端采集适配层，避免页面直接依赖具体分析平台。

#### 前端步骤

1. 新增 `frontend/src/analytics/index.ts`。
2. 定义 `track(eventName, properties)`。
3. 定义事件类型，禁止页面自由拼写事件名。
4. 自动附加：
   - 当前路由；
   - 用户 ID；
   - 会话 ID；
   - 客户端时间；
   - 应用版本。
5. 在本地开发环境打印事件，在生产环境发送到配置的采集端。

#### 后端步骤

可选 MVP：新增内部事件接口和表；若后续接入第三方平台，保留相同前端接口。

- `POST /api/v1/events`。
- 表 `analytics_events`：`user_id`、`event_name`、`properties`、`session_id`、`occurred_at`。
- 限制事件体积和允许的事件名。

#### 首批事件

- `register_completed`。
- `record_page_viewed`。
- `record_mode_switched`。
- `record_extras_expanded`。
- `record_submit_clicked`。
- `record_created`。
- `record_failed`。

#### 验收标准

- 同一次行为不会重复上报。
- 上报失败不阻塞主流程。
- 不上传笔记、邮箱、图片 URL 等敏感内容。
- 可以按用户、日期和事件名查询。

### TASK-003：建立自动化测试基线

- 状态：`TODO`
- 优先级：P0
- 角色：FE、BE、QA
- 工作量：M
- 依赖：无

#### 后端步骤

1. 为 `CoffeeLogService` 增加单元测试。
2. 为 `StatsService` 增加空数据、少量数据和有效数据测试。
3. 为记录创建接口增加鉴权、校验和空字段测试。
4. 在 CI 中执行 `go test ./...`。

#### 前端步骤

1. 引入 Vitest 和 Vue Test Utils。
2. 为 DTO 转换、快捷记录载荷和统计阈值增加单元测试。
3. 引入 Playwright 或等价工具覆盖核心 E2E：注册、首次记录、再次冲煮。
4. 在 CI 中执行类型检查和生产构建。

#### 验收标准

- 核心记录流程存在至少一条端到端测试。
- 统计空值规则存在自动化测试。
- CI 失败会阻止发布。

---

## 5. Sprint 1：数据可信与统计治理

### DATA-001：设计记录数据来源与有效性字段

- 状态：`DONE` ✅
- 优先级：P0
- 角色：PM、BE、FE
- 工作量：M
- 依赖：TASK-001
- 完成日期：2026-08-18

#### 推荐数据库字段

在 `coffee_logs` 增加：

| 字段 | 类型 | 说明 |
|---|---|---|
| `record_mode` | varchar(20) | `quick` 或 `detailed` |
| `coffee_name_source` | varchar(30) | `user_input` / `system_suggested` |
| `notes_source` | varchar(30) | `user_input` / `ai_generated` / `empty` |
| `shop_source` | varchar(30) | `user_input` / `recent_reuse` / `empty` |
| `sensory_recorded` | boolean | 六维感官是否由用户主动提交 |
| `source_log_id` | nullable bigint | 再次冲煮来源记录 |
| `is_test_data` | boolean | 测试数据标记 |

#### 设计决策

- MVP 使用 `sensory_recorded` 区分默认 `0` 与真实感官数据，降低一次性改为可空字段的风险。
- 后续可将六维字段改为可空整数；不建议同一个版本同时改类型和统计逻辑。
- `AI Summary` 与用户 `Notes` 必须继续分开保存。

#### 建议文件

- `backend/internal/model/coffee_log.go`。
- `backend/cmd/server/main.go` 的 `AutoMigrate` 清单。
- `backend/internal/service/coffee_log_service.go`。
- `frontend/src/api/coffeeLog.ts`。
- `frontend/src/stores/coffeeLog.ts`。

#### 验收标准

- 新记录能明确知道创建模式和感官数据是否有效。
- 再次冲煮记录能追溯来源。
- 旧客户端缺少新字段时仍可兼容。

### DATA-002：调整创建/更新记录接口契约

- 状态：`DONE` ✅
- 优先级：P0
- 角色：BE、FE
- 工作量：M
- 依赖：DATA-001
- 完成日期：2026-08-18

#### 后端步骤

1. 扩展 `CreateCoffeeLogRequest`：
   - `record_mode`；
   - `sensory_recorded`；
   - `source_log_id`；
   - 各字段 source。
2. `record_mode=quick` 时：
   - 允许空 `shop_name`；
   - 允许空 `notes`；
   - 感官分数不参与统计；
   - 允许服务端根据类型生成展示名称，但标记为 `system_suggested`。
3. `record_mode=detailed` 且 `sensory_recorded=true` 时，校验六维值均在 0—5。
4. 更新接口允许清空字符串，不能继续以“空字符串代表不更新”。

#### 重要技术修正

当前 `UpdateCoffeeLogRequest` 对字符串使用值类型，服务层又通过 `!= ""` 判断更新，导致无法主动清空门店、笔记等字段。建议改为 `*string`。

#### 前端步骤

1. 更新 `CreateCoffeeLogParams` 和 `CoffeeLogDTO`。
2. `toCreatePayload` 仅提交用户真实填写的数据。
3. 快捷与精细模式分别构建明确载荷，不再共享一组虚假默认值。
4. 增加 payload 单元测试。

#### 验收标准

- 快捷记录可以不提交门店、笔记和感官分数。
- 更新接口可以主动清空可选字段。
- 非法感官数据返回可理解的业务错误。

### DATA-003：停止生成虚假默认数据

- 状态：`DONE` ✅
- 优先级：P0
- 角色：FE、BE
- 工作量：S
- 依赖：DATA-002
- 完成日期：2026-08-18

#### 前端步骤

更新 `frontend/src/pages/CreateCoffeeLog.vue`：

1. 删除 `shop_name: 'Local Coffee Spot'`。
2. 删除默认笔记“一杯温润安静的手账记录。”。
3. 快捷模式未选风味时设置 `sensory_recorded=false`。
4. 默认图片可以继续作为展示封面，但不得标记为用户上传。
5. 默认咖啡名称标记为系统建议，并允许详情页编辑。

#### 后端步骤

1. `ShopService.EnsureShopForLog` 仅在 `ShopName` 非空时执行。
2. 用户未开启 AI 时，不生成看似真实的 Mock 感官总结。
3. 无 AI 总结时详情页使用空状态，不自动填充文案。

#### 验收标准

- 创建一条最简快捷记录后，数据库中的门店、笔记和 AI 摘要为空。
- 最简记录不会创建名为 `Local Coffee Spot` 的门店。
- 最简记录不进入风味雷达平均值。

### DATA-004：识别并迁移历史污染数据

- 状态：`TODO`
- 优先级：P0
- 角色：BE、QA
- 工作量：M
- 依赖：DATA-001、DATA-003

#### 开发步骤

1. 编写只读审计命令，统计：
   - 默认门店记录数；
   - 默认笔记记录数；
   - 六维值完全等于快捷默认组合的记录数；
   - 自动生成门店数量。
2. 输出待迁移记录 ID 和用户数量。
3. 编写幂等迁移命令：
   - 默认门店改为空；
   - 默认笔记改为空；
   - 疑似默认感官数据标记 `sensory_recorded=false`；
   - 不直接删除用户可能真实创建的门店，先生成审计报告。
4. 执行前备份数据库。
5. 测试环境验证后再执行生产迁移。

#### 建议文件

- 新增：`backend/cmd/data_audit/main.go`。
- 新增：`backend/cmd/data_cleanup/main.go`。
- 新增：`backend/internal/migration/coffee_log_quality.go`。

#### 验收标准

- 命令重复执行不会二次破坏数据。
- 每次迁移输出处理数量和失败 ID。
- 可通过数据库备份回滚。

### STATS-001：统计仓储仅聚合有效数据

- 状态：`TODO`
- 优先级：P0
- 角色：BE
- 工作量：M
- 依赖：DATA-001、DATA-004

#### 开发步骤

1. 更新 `backend/internal/repository/stats_repository.go`。
2. 风味平均值增加 `sensory_recorded=true` 条件。
3. 咖啡类型统计排除测试数据和无效记录。
4. 门店和标签统计排除空值。
5. 人格规则在总有效记录不足 8 条时返回空数组和进度信息。
6. 为所有聚合函数增加空集合测试。

#### 验收标准

- 没有有效感官记录时不返回全 0 雷达结论。
- 3 条以下记录不生成确定性人格。
- 历史污染数据清理前后统计结果可解释。

### STATS-002：统计接口返回样本量与范围

- 状态：`DONE` ✅
- 优先级：P0
- 角色：BE、FE
- 工作量：M
- 依赖：STATS-001
- 完成日期：2026-08-18

#### API 变更

建议统一返回：

```json
{
  "data": {},
  "meta": {
    "sample_count": 12,
    "valid_sensory_count": 7,
    "date_from": "2026-07-01",
    "date_to": "2026-08-18",
    "threshold": 3,
    "is_ready": true
  }
}
```

#### 前端步骤

更新：

- `frontend/src/api/stats.ts`。
- `frontend/src/stores/coffeeLog.ts`。
- `frontend/src/pages/Stats.vue`。
- `frontend/src/pages/MonthlyReview.vue`。

展示规则：

- `is_ready=false` 时展示积累进度和记录入口。
- 雷达旁展示“基于 N 条有效感官记录”。
- 人格旁展示“再记录 N 杯后生成”。

#### 验收标准

- 所有统计结论均可看到样本量或生成门槛。
- 空数据不渲染误导性图表。
- 旧客户端在灰度期间可以兼容。

---

## 6. Sprint 2：首次激活与记录反馈

### ONBOARD-001：增加用户引导状态

- 状态：`DONE` ✅
- 优先级：P0
- 角色：BE、FE
- 工作量：M
- 依赖：DATA-002
- 完成日期：2026-08-19

#### 数据字段

在 `users` 增加：

- `onboarding_completed` boolean。
- `preferred_log_mode` varchar(20)。
- `preferred_coffee_types` varchar/json。
- `first_record_at` nullable datetime。

#### API

- 扩展 `GET /api/v1/users/me`。
- 扩展 `PUT /api/v1/users/me`。
- 或新增 `PUT /api/v1/users/me/onboarding`。

#### 验收标准

- 用户可以跳过引导。
- 已完成用户不会重复进入引导。
- 跨设备登录后引导状态一致。

### ONBOARD-002：实现三步首次引导页

- 状态：`DONE` ✅
- 优先级：P0
- 角色：UX、FE
- 工作量：M
- 依赖：ONBOARD-001
- 完成日期：2026-08-19

#### 页面步骤

1. 常喝类型多选。
2. 记录偏好：快捷或精细。
3. 引导记录第一杯。

#### 建议文件

- 新增：`frontend/src/pages/Onboarding.vue`。
- 新增：`frontend/src/stores/onboarding.ts`。
- 更新：`frontend/src/router/index.ts`。
- 更新：`frontend/src/pages/Register.vue`。
- 更新：`frontend/src/stores/auth.ts`。

#### 路由规则

- 注册成功且未完成引导 → `/onboarding`。
- 已登录但未完成引导访问首页 → 引导页。
- 用户选择跳过 → 首页，同时保留“记录第一杯”入口。

#### 埋点

- `onboarding_started`。
- `onboarding_step_completed`。
- `onboarding_skipped`。
- `onboarding_completed`。

#### 验收标准

- 30 秒内可以完成或跳过。
- 返回操作不会丢失已选择内容。
- 320px 宽度下无横向滚动。

### ONBOARD-003：第一杯专属流程

- 状态：`TODO`
- 优先级：P0
- 角色：PM、FE、BE
- 工作量：M
- 依赖：ONBOARD-002、DATA-003

#### 开发步骤

1. 引导页将偏好带入记录页，但只作为建议。
2. 记录页显示“记录你的第一杯”标题。
3. 第一条记录成功后更新 `first_record_at`。
4. 返回基础反馈卡，不直接跳到复杂统计页。
5. 首页空状态根据是否完成第一杯展示不同 CTA。

#### 验收标准

- 首杯流程不要求专业感官评分。
- 保存成功后明确告知后续洞察需要的记录数量。
- 首杯事件只触发一次。

### SUCCESS-001：增加记录进度接口

- 状态：`DONE` ✅
- 优先级：P0
- 角色：BE
- 工作量：S
- 依赖：STATS-001

#### 推荐 API

`GET /api/v1/stats/record-progress?record_id={id}`

返回：

- 本月第几杯。
- 总记录数。
- 是否为第一杯。
- 下一项洞察名称。
- 距离阈值还差多少条。
- 可展示的轻量结论；不足样本时为空。

#### 验收标准

- 只返回当前用户数据。
- 记录不存在时返回 404。
- 结论使用有效记录口径。

### SUCCESS-002：实现记录成功反馈页面

- 状态：`DONE` ✅
- 优先级：P0
- 角色：UX、FE
- 工作量：M
- 依赖：SUCCESS-001
- 完成日期：2026-08-18

#### 页面内容

- 本次图片、类型、心情、日期。
- 本月累计杯数。
- 下一项洞察进度。
- 查看详情。
- 返回首页。
- 再次冲煮。
- 分享；可放入二级操作。

#### 建议文件

- 新增：`frontend/src/pages/RecordSuccess.vue`。
- 更新：`frontend/src/router/index.ts`。
- 更新：`frontend/src/pages/CreateCoffeeLog.vue`。
- 新增：`frontend/src/api/progress.ts`，或并入 `stats.ts`。

#### 路由

保存成功后进入 `/coffee/:id/success`，刷新后仍可以恢复内容。

#### 埋点

- `record_success_viewed`。
- `record_success_action`，参数为 `detail/home/rebrew/share`。

#### 验收标准

- 保存成功后 300ms 内出现反馈状态。
- 进度接口失败时仍可展示基础成功信息。
- 页面刷新和浏览器返回行为正确。

---

## 7. Sprint 3：周度回顾、记忆与再次冲煮

### REVIEW-001：周度聚合接口

- 状态：`DONE` ✅
- 优先级：P1
- 角色：BE
- 工作量：M
- 依赖：STATS-001、STATS-002
- 完成日期：2026-08-19

#### API

`GET /api/v1/stats/weekly-review?week=2026-W34`

返回：

- 周起止日期。
- 有效记录数。
- 最常喝类型。
- 最常见心情和场景。
- 一条可信趋势。
- 一条历史回忆候选。
- 与前一周的变化；样本不足时为空。

#### 后端文件

- `backend/internal/repository/stats_repository.go`。
- `backend/internal/service/stats_service.go`。
- `backend/internal/handler/stats_handler.go`。
- `backend/internal/router/router.go`。

#### 验收标准

- 正确处理跨月和跨年周。
- 时区统一使用用户或系统配置时区。
- 0 条、1 条和多条数据均有测试。

### REVIEW-002：周度回顾页面与首页卡片

- 状态：`DONE` ✅
- 优先级：P1
- 角色：UX、FE
- 工作量：M
- 依赖：REVIEW-001
- 完成日期：2026-08-19

#### 页面

- 新增：`frontend/src/pages/WeeklyReview.vue`。
- 首页增加“本周咖啡足迹”入口。
- 咖迹页在月度回顾上方增加周度回顾。

#### 空状态

- 0 条：引导记录第一杯。
- 1—2 条：展示事实，不生成趋势。
- 达到阈值：展示完整周报。

#### 埋点

- `weekly_review_entry_viewed`。
- `review_viewed`，`period=weekly`。
- `review_action_clicked`。

### MEMORY-001：历史记忆接口

- 状态：`TODO`
- 优先级：P1
- 角色：BE
- 工作量：M
- 依赖：DATA-004

#### API

`GET /api/v1/coffee-logs/memories`

返回最多三条候选：

- 上月的今天。
- 去年的今天。
- 最近高频再次饮用。
- 最近被再次冲煮的来源记录。

#### 规则

- 不返回已删除记录。
- 同一天不重复同一记录。
- 没有精确日期时允许在 ±3 天窗口内匹配。

### MEMORY-002：首页记忆卡片

- 状态：`TODO`
- 优先级：P1
- 角色：FE、UX
- 工作量：S
- 依赖：MEMORY-001

#### 验收标准

- 每次首页最多出现一张记忆卡。
- 可以关闭当次卡片。
- 进入详情和再次冲煮入口明确。
- 不抢占今日记录主 CTA。

### REBREW-001：规范再次冲煮预填规则

- 状态：`DONE` ✅
- 优先级：P1
- 角色：FE、BE
- 工作量：M
- 依赖：DATA-001、DATA-002
- 完成日期：2026-08-19

#### 当前状态

`CreateCoffeeLog.vue` 已支持 `from_log_id`，但会带入心情、生活标签、感官评分和图片等较多信息。

#### 新规则

自动带入：

- 咖啡类型。
- 豆子。
- 门店。
- 冲煮参数。
- 可选封面。

仅作为参考、不直接提交：

- 六维感官评分。
- 风味标签。

不带入：

- 心情。
- 生活标签。
- 笔记。
- 日期。

#### 后端步骤

- 保存 `source_log_id`。
- 校验来源记录属于当前用户。
- 防止循环来源关系。

#### 前端步骤

- 展示“基于某次记录再次冲煮”。
- 感官步骤显示上次值的淡色参考线。
- 用户未主动修改时 `sensory_recorded=false`。

#### 验收标准

- 再次冲煮不会复制旧心情和旧笔记。
- 保存后可以追溯来源记录。

### REBREW-002：两次冲煮结果对比

- 状态：`TODO`
- 优先级：P1
- 角色：BE、FE
- 工作量：M
- 依赖：REBREW-001

#### API

可直接通过当前记录和 `source_log_id` 获取；若前端组合复杂，再新增：

`GET /api/v1/coffee-logs/:id/comparison`

#### 页面内容

- 参数差异。
- 六维感官差异。
- 风味标签变化。
- 笔记差异不自动评价。

#### 验收标准

- 只有双方都有有效感官评分时显示雷达差异。
- 空参数显示“未记录”，不显示为 0。

---

## 8. Sprint 4：门店、豆子与 AI 信任

### REUSE-001：最近使用门店和豆子

- 状态：`TODO`
- 优先级：P1
- 角色：BE、FE
- 工作量：M
- 依赖：DATA-003

#### API

- `GET /api/v1/coffee-shops/recent?limit=5`。
- `GET /api/v1/coffee-beans/recent?limit=5`。

#### 后端规则

- 按当前用户最近一条关联记录排序。
- 同一门店或豆子去重。
- 空门店不参与。

#### 前端步骤

- 精细记录页优先展示最近项。
- 选择后允许清除。
- 记录来源标记为 `recent_reuse`。

### REUSE-002：门店常点咖啡与时间线

- 状态：`TODO`
- 优先级：P1
- 角色：BE、FE
- 工作量：M
- 依赖：REUSE-001

#### 后端

扩展门店详情：

- 到访次数。
- 最近到访。
- 常喝类型 Top 3。
- 关联记录按月分组。

#### 前端

更新 `CoffeeShopDetail.vue`：

- 顶部展示到访摘要。
- 关联记录改为紧凑时间线。
- 增加“在这里再记一杯”入口。

### REUSE-003：咖啡豆赏味期与冲煮历史

- 状态：`TODO`
- 优先级：P1
- 角色：BE、FE
- 工作量：L
- 依赖：REUSE-001、REBREW-001

#### 数据字段

在 `coffee_beans` 增加：

- `opened_at`。
- `roasted_at`。
- `best_before`。

#### 页面能力

- 豆子详情页。
- 当前赏味状态。
- 冲煮次数。
- 常用参数。
- 有效感官平均值。
- 最佳记录和再次冲煮入口。

### AI-001：AI 数据来源说明

- 状态：`TODO`
- 优先级：P1
- 角色：PM、FE、BE
- 工作量：M
- 依赖：DATA-002、STATS-002

#### 前端

- AI 开关旁展示“将发送哪些字段”。
- 提供隐私说明弹层。
- 生成结果展示“基于 N 条记录”。
- 标记 Mock 模式与外部模型模式。

#### 后端

- AI 响应增加：`provider`、`model`、`sample_count`、`generated_at`。
- 日志不得记录完整用户笔记和密钥。
- 用户关闭 AI 时不调用外部服务。

#### 验收标准

- 用户可以在提交前理解数据流向。
- 关闭 AI 后网络请求中不存在 AI 调用。

### AI-002：AI 结果反馈

- 状态：`TODO`
- 优先级：P1
- 角色：BE、FE
- 工作量：M
- 依赖：AI-001、TASK-002

#### 数据表

`ai_feedback`：

- `user_id`。
- `content_type`。
- `content_id`。
- `feedback_type`。
- `model`。
- `created_at`。

#### API

`POST /api/v1/ai/feedback`

#### 反馈项

- 有帮助。
- 不准确。
- 太笼统。
- 不符合表达。

#### 验收标准

- 同一内容重复反馈时更新而不是新增重复行。
- 反馈失败不影响查看内容。

---

## 9. Sprint 5：分享、性能与发布质量

### SHARE-001：分享卡内容层级优化

- 状态：`PARTIAL`
- 优先级：P2
- 角色：UX、FE
- 工作量：M
- 依赖：STATS-002

#### 开发步骤

1. 为单杯、周度、月度和人格建立统一模板规范。
2. 分享卡只展示用户真实数据。
3. 缺失字段自动收起，不显示空占位。
4. 保留产品标识但不遮挡内容。
5. 导出前检查跨域图片和字体加载状态。

#### 验收标准

- 3:4、1:1、9:16 三种比例可正确导出。
- Android、iOS 和桌面浏览器输出一致。
- 无数据字段不会显示 `0` 或 `undefined`。

### SHARE-002：公开只读分享链接 MVP

- 状态：`TODO`
- 优先级：P2
- 角色：PM、BE、FE
- 工作量：L
- 依赖：SHARE-001、TASK-002

#### 数据表

`share_links`：

- `user_id`。
- `resource_type`。
- `resource_id`。
- `token_hash`。
- `expires_at`。
- `revoked_at`。
- `created_at`。

#### API

- `POST /api/v1/share-links`。
- `DELETE /api/v1/share-links/:id`。
- `GET /api/v1/public/share/:token`，公开只读。

#### 安全要求

- Token 使用高熵随机值，数据库仅保存 Hash。
- 分享页不得暴露邮箱、用户 ID 或私人笔记；私人笔记需单独授权。
- 用户可以随时撤销。
- 接口增加限流。

#### 前端

- 新增公开分享路由。
- 分享页不依赖登录态。
- 增加“创建我的咖啡手账”入口并附带来源参数。

### PERF-001：路由懒加载与包体优化

- 状态：`DONE` ✅
- 优先级：P1
- 角色：FE
- 工作量：M
- 依赖：无
- 完成日期：2026-08-19

#### 当前问题

生产构建主 JS 包已超过 500 kB 警告线，所有页面在路由文件中同步导入。

#### 开发步骤

1. 首页和启动页可同步加载，其余页面改为动态 import。
2. 将 `html2canvas` 仅在分享导出时加载。
3. 将图表相关代码按页面拆包。
4. 检查 Google Fonts 和 Unsplash 外部资源失败时的降级。
5. 图片列表使用懒加载和合适尺寸。

#### 验收标准

- 首屏主包明显下降。
- 分享、统计和记录页面形成独立 chunk。
- 慢速网络下仍显示可用骨架和系统字体。

### A11Y-001：无障碍与触控验收

- 状态：`PARTIAL`
- 优先级：P1
- 角色：FE、QA、UX
- 工作量：M
- 依赖：主要页面开发完成

#### 检查项

- 图标按钮有 `aria-label`。
- 可点击区域不小于 36×36px，核心操作目标 44×44px。
- 表单错误可被读屏读取。
- 颜色对比不只依赖颜色区分状态。
- 键盘可以完成注册、记录和回顾。
- 动画遵循 `prefers-reduced-motion`。
- 图片具有可理解的替代文本。

### QA-001：全链路回归测试

- 状态：`TODO`
- 优先级：P0
- 角色：QA、FE、BE
- 工作量：L
- 依赖：所有计划上线任务

#### 必测场景

1. 新用户注册 → 引导 → 第一杯 → 成功反馈。
2. 老用户跳过引导。
3. 最简快捷记录不产生虚假数据。
4. 精细记录完整提交六维感官。
5. 修改并清空门店、笔记、标签。
6. 再次冲煮不复制心情和笔记。
7. 统计样本不足和样本充足状态。
8. 周/月回顾跨月、跨年。
9. AI 开启、关闭、失败和 Mock 状态。
10. 分享导出和公开链接撤销。
11. 401、网络超时、重复提交和图片上传失败。
12. 320、375、430px 宽度及常见桌面尺寸。

#### 发布门槛

- P0/P1 缺陷为 0。
- P2 缺陷有明确处理结论。
- 核心 E2E 全部通过。
- 数据迁移已在测试库演练。

---

## 10. 前后端接口开发顺序

以下顺序用于减少联调阻塞：

1. 冻结字段和指标定义。
2. 完成数据库字段和 Go Model。
3. 完成请求/响应 DTO 与接口文档。
4. 前端使用静态 Mock 数据开发页面。
5. 后端完成 Repository 和 Service。
6. 后端完成 Handler、Router 和单元测试。
7. 前端更新 API 类型与 Store。
8. 完成联调和错误状态。
9. 完成埋点。
10. 完成 E2E、迁移演练和灰度发布。

接口变更必须先更新文档或 OpenAPI 契约，再修改前端调用。

---

## 11. 数据库变更总表

### 11.1 `users`

| 字段 | Sprint | 用途 |
|---|---:|---|
| `onboarding_completed` | 2 | 引导状态 |
| `preferred_log_mode` | 2 | 默认记录方式 |
| `preferred_coffee_types` | 2 | 首次偏好 |
| `first_record_at` | 2 | 激活时间 |

### 11.2 `coffee_logs`

| 字段 | Sprint | 用途 |
|---|---:|---|
| `record_mode` | 1 | 快捷/精细来源 |
| `coffee_name_source` | 1 | 名称数据来源 |
| `notes_source` | 1 | 笔记数据来源 |
| `shop_source` | 1 | 门店数据来源 |
| `sensory_recorded` | 1 | 感官有效性 |
| `source_log_id` | 1/3 | 再次冲煮来源 |
| `is_test_data` | 1 | 测试数据隔离 |

### 11.3 `coffee_beans`

| 字段 | Sprint | 用途 |
|---|---:|---|
| `opened_at` | 4 | 开封日期 |
| `roasted_at` | 4 | 烘焙日期 |
| `best_before` | 4 | 赏味期 |

### 11.4 新表

| 表 | Sprint | 用途 |
|---|---:|---|
| `analytics_events` | 0 | 产品事件；若采用外部平台可不建 |
| `ai_feedback` | 4 | AI 内容反馈 |
| `share_links` | 5 | 公开分享链接 |

### 11.5 迁移要求

- GORM `AutoMigrate` 仅负责新增结构，不负责业务数据清理。
- 所有历史回填必须使用独立、幂等、可审计命令。
- 生产执行前必须备份。
- 大表更新应分批执行并记录游标。

---

## 12. API 变更总表

| 方法 | 路径 | Sprint | 说明 |
|---|---|---:|---|
| PUT | `/users/me/onboarding` | 2 | 保存引导状态与偏好 |
| GET | `/stats/record-progress` | 2 | 记录成功进度 |
| GET | `/stats/weekly-review` | 3 | 周度回顾 |
| GET | `/coffee-logs/memories` | 3 | 历史记忆 |
| GET | `/coffee-logs/:id/comparison` | 3 | 再次冲煮对比，可选 |
| GET | `/coffee-shops/recent` | 4 | 最近使用门店 |
| GET | `/coffee-beans/recent` | 4 | 最近使用豆子 |
| POST | `/ai/feedback` | 4 | AI 反馈 |
| POST | `/share-links` | 5 | 创建分享链接 |
| DELETE | `/share-links/:id` | 5 | 撤销分享链接 |
| GET | `/public/share/:token` | 5 | 公开只读分享 |

现有接口需扩展：

- `POST /coffee-logs`：数据来源、记录模式、感官有效性、来源记录。
- `PUT /coffee-logs/:id`：支持清空可选字段。
- `/stats/*`：增加 `meta`、阈值和有效样本数。
- `/ai/*`：增加模型、样本数和生成时间。

---

## 13. 前端页面与组件清单

### 新增页面

- `Onboarding.vue`。
- `RecordSuccess.vue`。
- `WeeklyReview.vue`。
- 可选：`CoffeeBeanDetail.vue`。
- 可选：`PublicShare.vue`。

### 重点修改页面

- `Register.vue`：注册后跳转引导。
- `Home.vue`：首次记录、周报和记忆入口。
- `CreateCoffeeLog.vue`：真实数据载荷、再次冲煮规则。
- `CoffeeDetail.vue`：空数据、来源记录、对比入口。
- `Stats.vue`：样本量、阈值和分层结构。
- `MonthlyReview.vue`：有效样本和时间范围。
- `CoffeeShopDetail.vue`：常点咖啡与再记一杯。
- `ShareCard.vue`：真实数据和动态空字段。
- `Profile.vue`：隐私、导出、提醒和 AI 设置。

### 建议新增组件

- `InsightProgress.vue`：洞察生成进度。
- `SampleMeta.vue`：样本量与统计范围。
- `RecordSuccessCard.vue`：成功反馈主体。
- `WeeklyReviewCard.vue`：周报入口。
- `MemoryCard.vue`：历史记忆。
- `RecentSelector.vue`：最近门店/豆子。
- `AIDataDisclosure.vue`：AI 数据说明。
- `EmptyState.vue`：统一空状态。

---

## 14. 验收用例矩阵

| 场景 | 数据状态 | 预期结果 |
|---|---|---|
| 最简快捷记录 | 只有类型、心情、默认封面 | 保存成功；门店/笔记为空；无有效感官 |
| 快捷记录 + 风味预设 | 用户主动选择风味 | 可按产品决策标记有效感官，必须明确是用户确认 |
| 精细记录 | 完整六维评分 | 进入雷达和统计 |
| 空门店 | 未选择门店 | 不创建门店档案 |
| 清空门店 | 旧记录有门店 | 更新后成功置空 |
| 样本不足 | 2 条记录 | 不生成人格，展示进度 |
| 无有效感官 | 多条记录但未评分 | 不显示全 0 雷达 |
| 再次冲煮 | 来源记录含心情/笔记 | 不复制心情和笔记 |
| AI 关闭 | 用户关闭 AI | 不调用外部服务，不生成假 AI 文案 |
| AI 失败 | 外部服务超时 | 记录仍保存，AI 状态可重试 |
| 分享缺字段 | 无门店/笔记 | 卡片自动重排，无空占位 |
| 迁移重复执行 | 已迁移数据 | 不重复修改，不报破坏性错误 |

---

## 15. 每个 Sprint 的发布检查

### 发布前

- [ ] 需求、接口和字段已冻结。
- [ ] 数据库备份完成。
- [ ] 迁移命令已在测试库演练。
- [ ] 前后端自动化测试通过。
- [ ] 生产构建通过。
- [ ] 核心页面完成手机尺寸检查。
- [ ] 埋点在测试环境可以查询。
- [ ] 功能开关和回滚方案可用。

### 灰度中

- [ ] 监控创建记录错误率。
- [ ] 监控 API P95 响应时间。
- [ ] 检查虚假默认数据是否继续增长。
- [ ] 检查首次记录和保存完成率。
- [ ] 检查前端异常和白屏率。

### 全量后

- [ ] 24 小时内复盘错误和漏斗。
- [ ] 7 天后评估激活与留存。
- [ ] 对比版本前后的数据质量。
- [ ] 记录未达预期的假设和下一步实验。

---

## 16. 首个可执行开发批次

建议下一次实际开发按以下顺序开始，不并行跳过数据基础：

1. `TASK-001`：冻结有效记录与有效感官口径。
2. `DATA-001`：新增数据来源和有效性字段。
3. `DATA-002`：调整创建/更新接口契约。
4. `DATA-003`：清除新记录的虚假默认数据。
5. `DATA-004`：审计并迁移历史数据。
6. `STATS-001`：统计只使用有效数据。
7. `STATS-002`：接口和页面增加样本量。
8. `SUCCESS-001/002`：实现记录成功反馈。
9. `ONBOARD-001/002/003`：实现首次激活流程。

首批完成后再进入周度回顾和再次冲煮，可以避免基于不可信数据建设新的洞察能力。

---

## 17. 推荐的首批开发 PR 拆分

为降低审查和回滚风险，建议拆成以下独立 PR：

1. **PR-01 数据字段与兼容读取**  
   只新增字段、DTO 和兼容读取，不改变线上行为。

2. **PR-02 记录载荷 V2**  
   前端发送记录模式、来源和感官有效性；服务端校验。

3. **PR-03 停止虚假默认值**  
   删除默认门店、笔记、分数和无条件 Mock AI 文案。

4. **PR-04 数据审计与迁移工具**  
   提供只读报告和幂等清理命令。

5. **PR-05 可信统计 V2**  
   聚合过滤、样本量和阈值展示。

6. **PR-06 记录成功反馈**  
   进度接口、成功页和埋点。

7. **PR-07 首次引导**  
   用户字段、引导页、注册跳转和第一杯流程。

每个 PR 必须可独立回滚，不应同时包含无关视觉重构。

---

## 18. 最终里程碑验收

### 里程碑 A：数据可信

- 新记录不再产生虚假门店和笔记。
- 未评分记录不进入感官统计。
- 统计结论显示样本量和时间范围。
- 历史污染数据完成审计和迁移。

### 里程碑 B：首次激活

- 注册后 30 秒内可完成引导。
- 用户可在 60 秒内完成第一杯。
- 保存后展示明确反馈和下一步。

### 里程碑 C：留存闭环

- 用户能够查看周度回顾。
- 首页能够唤醒一条历史记忆。
- 再次冲煮不会复制瞬时数据。

### 里程碑 D：复用与信任

- 最近门店和豆子可以快速复用。
- AI 明确说明数据来源和模型状态。
- 用户可以对 AI 结果反馈。

### 里程碑 E：增长验证

- 分享卡只展示真实数据。
- 分享链接可以撤销和归因。
- 性能、无障碍和核心 E2E 达到发布标准。

