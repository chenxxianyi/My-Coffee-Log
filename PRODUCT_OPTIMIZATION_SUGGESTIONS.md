# My Coffee Log 产品优化拆分步骤

## 0. 产品优化总方向

My Coffee Log 当前最适合的方向不是“专业咖啡参数工具”，而是一个高审美、低负担、可分享的咖啡生活方式手账产品。

核心价值：

> 让用户用很低成本记录一杯咖啡，并把它变成一段好看的生活记忆。

建议产品定位：

> 把每天的一杯咖啡，记录成属于自己的生活方式手账。

英文表达：

> My Coffee Log helps coffee lovers turn everyday cups into beautiful personal memories.

---

## 1. 优化步骤总览

| 步骤 | 优化点 | 优先级 | 核心目标 | 状态 |
|------|--------|--------|----------|------|
| Step 1 | Quick Log 快速记录模式 | P0 | 降低首次记录成本 | ✅ 已完成 |
| Step 2 | 首页咖啡生活仪表盘 | P0 | 提升首页情绪价值和使用效率 | ✅ 已完成 |
| Step 3 | 记录完成后的即时正反馈 | P0 | 让用户保存后立刻获得奖励 | ✅ 已完成 |
| Step 4 | 咖啡详情页视觉强化 | P0 | 让每条记录像一页生活方式杂志 | ✅ 已完成 |
| Step 5 | AI 文案风格升级 | P0 | 从风味分析升级为 editorial 文案 | ✅ 已完成 |
| Step 6 | 分享卡片模板优化 | P0 | 提升外部传播能力 | ✅ 已完成 |
| Step 7 | 咖啡心情日记 | P1 | 让记录从咖啡扩展到生活状态 | ✅ 已完成 |
| Step 8 | 统计页升级为个人咖啡画像 | P1 | 提升统计页情绪价值和分享价值 | ✅ 已完成 |
| Step 9 | 复刻上一杯 | P1 | 提高重复记录效率 | ✅ 已完成 |
| Step 10 | 月度咖啡回顾 | P1 | 提升留存和月度分享 | ✅ 已完成 |
| Step 11 | 咖啡店收藏 | P1 | 沉淀用户长期消费场景 | ✅ 已完成 |
| Step 12 | 咖啡豆档案 | P2 | 服务深度手冲用户 | ✅ 已完成 |
| Step 13 | 冲煮参数记录 | P2 | 扩展专业记录能力 | |
| Step 14 | 咖啡地图 | P2 | 扩展城市咖啡生活场景 | |
| Step 15 | 真实 AI 接入 | P2 | 提升文案与洞察个性化能力 | ✅ 已完成 |
| Step 16 | 轻量成就系统 | P2 | 增强长期记录动力 | |

---

## Step 1：Quick Log 快速记录模式 ✅ 已完成

### 优先级

P0

### 优化目标

把"记录一杯咖啡"的流程压缩到 3 秒左右，让用户可以低负担完成第一条记录。

### 核心改动

- 在创建页增加快速记录模式。
- 快速记录只保留照片、咖啡类型、心情、一键保存。
- 原有完整表单作为精细记录模式保留。
- 未填写字段使用合理默认值。

### 实现详情

- 创建页顶部增加 Quick / Detailed 模式切换按钮（pill toggle 风格）。
- Quick 模式为单页表单：封面图选择 + 咖啡类型选择 + 心情选择 + 咖啡名称（选填）。
- Quick 模式默认进入，点击"一键保存"直接提交。
- 未填写字段默认值：coffee_name 根据咖啡类型自动生成（如"手冲咖啡"），shop_name 默认 "Local Coffee Spot"，notes 默认 "一杯温润安静的手账记录。"，风味参数默认 3，flavor_tags 默认空。
- Detailed 模式保留原有 3 步向导流程不变。
- 文件上传根据当前模式写入对应表单。
- 保存成功后跳转详情页。

### 涉及模块

- 创建页 (`CreateCoffeeLog.vue`)
- Coffee Log Store
- Coffee Log API
- 后端 Coffee Log 创建接口

### 验收标准

- 用户可以不填写复杂参数直接保存一条咖啡记录。
- 快速记录保存成功后进入详情页。
- 精细记录模式仍然保留完整表单能力。

---

## Step 2：首页咖啡生活仪表盘 ✅ 已完成

### 优先级

P0

### 优化目标

让首页从普通功能入口升级为“咖啡生活仪表盘”，提升打开 App 后的情绪价值和使用效率。

### 核心改动

- 首页展示今日咖啡记录入口。
- 展示最近一杯咖啡。
- 展示本月咖啡记录数。
- 展示最常喝的咖啡类型。
- 展示最近偏好的风味标签。
- 展示一句 AI 生活方式文案。

### 实现详情

- 首页新增三种状态展示：今日已记录（绿色脉冲指示器 + 记录详情卡片）、最近记录 + Quick Log 提示、空状态欢迎卡片。
- 新增 `todayLog` 计算属性，自动匹配当日咖啡记录。
- 标题改为动态问候语（根据时段显示早安/午后/晚安等）。
- AI 生活文案从硬编码改为调用后端 `/api/v1/ai/lifestyle-quote` 接口生成，文案根据月度冲煮数、偏好咖啡类型、风味标签和最近心情动态组合。
- 月度概览从 3 列改为 2 列布局，加大数字字号，新增“偏好风味图谱”区块展示 Top 5 风味标签（带杯数统计）。
- 最近手账列表从 2 条扩展为 3 条。
- 后端新增 `GetRecentFlavorTags` 查询，返回用户 Top N 风味标签及使用次数。
- 后端 Stats Overview 响应新增 `recent_flavor_tags` 字段。
- 后端新增 `GenerateLifestyleQuote` 方法，根据月度统计、偏好类型、风味标签和心情生成编辑风格式生活文案。
- 后端新增 `/api/v1/ai/lifestyle-quote` POST 路由和 handler。
- 前端 Store 新增 `lifestyleQuote`、`recentFlavorTags`、`todayLog` 状态及 `fetchLifestyleQuote` action。
- 前端 API 层新增 `FlavorTagItem` 接口、`LifestyleQuoteRequest/Response` 接口和 `getLifestyleQuote` 函数。

### 涉及模块

- 首页 (`Home.vue`)
- Coffee Log Store (`coffeeLog.ts`)
- Stats API (`stats.ts`)
- AI 文案生成逻辑 (`ai_service.go`, `ai_handler.go`)
- Stats Repository (`stats_repository.go`)
- Stats Service (`stats_service.go`)
- Router (`router.go`)

### 验收标准

- ✅ 首页首屏能看到明确的记录入口（今日已记录/Quick Log 提示/空状态引导）。
- ✅ 首页能展示最近一条咖啡记录。
- ✅ 首页能展示本月记录数和偏好信息（月度冲煮数、最常喝类型、偏好风味图谱）。
- ✅ 首页展示 AI 生活方式文案（由后端动态生成）。
- ✅ 首页视觉符合 Nordic Minimal + Editorial 风格。

---

## Step 3：记录完成后的即时正反馈 ✅ 已完成

### 优先级

P0

### 优化目标

让用户完成记录后立刻获得奖励，增强继续使用和分享的动力。

### 核心改动

- 创建成功后自动进入详情页。
- 详情页立即展示 AI 文案。
- 展示风味雷达图。
- 展示分享卡片入口。
- 展示本月记录进度提示。

### 实现详情

- 创建页保存成功后携带 `?just_created=true` 查询参数跳转详情页。
- 创建页保存成功后异步刷新统计数据（`store.fetchStats()`），确保详情页月度进度数据实时更新。
- 详情页检测 `just_created` 状态，触发即时正反馈体验：
  - **庆祝横幅**：顶部滑入「手账记录成功」Banner，显示绿色勾选图标 + 本月第 N 杯计数，3 秒后自动淡出。
  - **庆祝粒子**：8 颗咖啡色系圆点粒子从底部向上飘升，配合 Banner 同步展示与消失。
  - **月度进度卡片**：在详情内容区展示「本月咖啡进度」卡片，包含渐变进度条（latte→terracotta）、里程碑标签（3/5/10/15/20 杯）、editorial 风格里程碑文案。
  - **分享 CTA 卡片**：深色底分享引导区块，突出「生成分享海报」按钮，鼓励用户即时分享。
  - **再记一杯按钮**：在操作栏替换原「生成分享海报」为「再记一杯」，引导连续记录。
  - **入场动画**：封面图 fade-in、标题区 slide-up、AI 评语延迟 0.15s、雷达图延迟 0.3s、标签区延迟 0.4s，形成交错式入场节奏。
- 非创建状态进入详情页时，不显示庆祝效果，保持原有静态展示。

### 涉及模块

- 创建页 (`CreateCoffeeLog.vue`)
- 咖啡详情页 (`CoffeeDetail.vue`)
- Coffee Log Store (`coffeeLog.ts`)
- Stats API (`stats.ts`)

### 验收标准

- ✅ 创建成功后自动跳转详情页。
- ✅ 详情页展示 AI 文案和风味雷达图。
- ✅ 用户可以从详情页直接进入分享卡片。
- ✅ 创建后详情页展示庆祝横幅和粒子动画。
- ✅ 创建后详情页展示本月咖啡进度卡片（含进度条和里程碑文案）。
- ✅ 创建后详情页展示分享 CTA 和「再记一杯」按钮。
- ✅ 创建后详情页各区域有交错式入场动画。
- ✅ 非创建状态进入详情页无庆祝效果，保持原有展示。

---

## Step 4：咖啡详情页视觉强化 ✅ 已完成

### 优先级

P0

### 优化目标

让每条咖啡记录都像一页精致的生活方式杂志，而不是普通数据详情页。

### 核心改动

- 强化大图展示。
- 使用杂志式标题排版。
- 突出日期、地点、心情、风味标签。
- 强化 AI editorial 文案展示。
- 优化风味雷达图展示位置。
- 增加明显的分享按钮。

### 实现详情

- **封面大图强化**：高度从 h-72 提升至 h-[340px]，增加 scale-105 微放大效果；多层渐变叠加（底部 warmWhite 渐变 + 顶部黑色渐变）营造杂志深度感；增加 paper grain 纸质纹理叠加；增加左上/右上角 bracket 装饰元素。
- **咖啡类型徽章**：从单行 badge 升级为双语展示（英文 + `/ 中文`），使用 espresso/90 底色 + backdrop-blur。
- **杂志式标题排版**：标题字号从 38px 提升至 42px，增加 tracking-wide 和 leading-[1.05]；日期两侧增加渐变 flanking lines 装饰；meta 信息条增加 MapPin 图标 + 竖线分隔 + moodLabel 中文标签。
- **Hairline 分隔线**：从单线升级为三段式（短灰线 + 长奶油线 + 短灰线），增加编辑感。
- **AI Editorial 文案强化**：增加 72px 装饰性大引号 `&ldquo;`；标题改为「AI 感官评语 / Editorial」双语；正文字号从 13.5px 提升至 15px，行高从 leading-relaxed 提升至 leading-[1.8]；增加 closing ornament（线 + 圆点 + 线）装饰；背景改为 bg-coffee-cream/60。
- **风味雷达图优化**：从左侧 2/5 布局改为居中展示；尺寸从 110px 提升至 150px；外框改为 170px 圆形容器（bg-coffee-cream/30 + border-coffee-latte/25）；维度标签改为英文缩写（Acid/Bitter/Sweet/Body/Aroma/After）；分数列表从右侧列表改为雷达下方 3 列网格，双语标签 + 简洁数字。
- **风味标签强化**：标题增加两侧 flanking lines；标签样式从 bg-coffee-cream 升级为 bg-coffee-warmWhite + shadow-sm，增加 tracking-wider。
- **味觉日记强化**：标题增加两侧 flanking lines；正文增加左侧 border-l-2 border-coffee-latte/30 引用线装饰；字号从 text-xs 提升至 text-[13px]，行高 leading-[1.75]。
- **浮动分享按钮**：新增右下角固定 FAB 按钮（terracotta 渐变 + 4px ring + hover:scale-110 + active:scale-95），底部标注「分享」文字。
- **底部操作栏重构**：分享按钮升级为全宽主按钮（py-3.5 + Share2 图标）；新增「再记一杯」按钮始终可见；删除按钮降级为次要操作。

### 涉及模块

- 咖啡详情页 (`CoffeeDetail.vue`)
- 风味雷达图组件 (`FlavorRadarChart.vue`)
- 咖啡常量 (`coffee.ts`) — 新增 `coffeeTypeShortLabel`, `moodLabel` 引用

### 验收标准

- ✅ 详情页首屏具备明显视觉吸引力（340px 大图 + 多层渐变 + 角括号装饰）。
- ✅ 用户能一眼看到照片、咖啡名称、AI 文案和核心风味。
- ✅ 详情页可以自然承接分享卡片功能（浮动 FAB + 底部全宽分享按钮）。
- ✅ 杂志式标题排版（42px serif italic + flanking lines + 双语 meta 信息）。
- ✅ AI editorial 文案具备杂志感（装饰性大引号 + closing ornament + 15px 正文）。
- ✅ 风味雷达图居中展示，尺寸更大，分数列表清晰双语。
- ✅ 分享入口明显且多入口（浮动 FAB + 顶栏按钮 + 底部主按钮）。

---

## Step 5：AI 文案风格升级 ✅ 已完成

### 优先级

P0

### 优化目标

让 AI 从“参数分析器”升级为“生活方式杂志编辑”，增强产品调性和分享价值。

### 核心改动

- 调整 AI 文案生成风格。
- 从风味参数描述升级为 editorial 文案。
- 为不同场景生成不同类型文案。
- 详情页使用风味摘要。
- 分享卡片使用生活方式文案。
- 统计页和月报使用偏好洞察文案。

### 文案示例

普通风味描述：

> 这杯咖啡酸度较高，带有柑橘和花香。

更适合产品调性的表达：

> 这是一杯明亮又安静的咖啡，像清晨窗边透进来的柑橘光，尾韵轻柔而干净。

英文示例：

> A bright and gentle cup, like a quiet morning with citrus light and a soft floral finish.

### 实现详情

- **DeepSeek API 客户端**：新增 `callDeepSeekAPI()` 函数，通过 OpenAI 兼容接口调用 DeepSeek Chat Completions API，支持 `OPENAI_API_KEY`、`OPENAI_BASE_URL`、`OPENAI_MODEL` 三个环境变量配置，默认模型 `deepseek-chat`，Temperature 0.85，MaxTokens 512。
- **Editorial 风格 System Prompt**：为三种场景分别设计了 editorial 风格的 system prompt，参考 Nordic minimal、日式手账、Kinfolk 杂志风格，要求用画面感、比喻和情绪传达风味，不直接罗列参数数值。
- **风味感官评语升级**：`GenerateFlavorSummary()` 优先调用 DeepSeek API 生成 editorial 风格评语，API 不可用时自动 fallback 到 `generateMockSummary()` 模板文案。请求结构新增 `coffee_name`、`mood`、`notes` 字段，为 AI 提供更丰富的上下文。
- **生活方式文案升级**：`GenerateLifestyleQuote()` 优先调用 DeepSeek API，API 不可用时 fallback 到 `generateMockLifestyleQuote()` 模板文案。
- **新增偏好洞察接口**：`POST /ai/flavor-insight` — 根据用户月度统计数据（杯数、偏好类型、风味标签、平均酸度/醇厚/甜感、最近心情）生成偏好洞察文案，用于统计页和月报场景。`GenerateFlavorInsightForUser()` 自动从 StatsService 获取 FlavorProfile 数据。
- **CoffeeLogService 联动**：创建和更新咖啡记录时，改用新的 `GenerateFlavorSummary()` 接口，传入完整上下文（咖啡名、心情、笔记），AI 评语质量显著提升。
- **优雅降级**：所有 AI 接口均实现 DeepSeek API → Mock 模板的双层 fallback，确保 API 不可用时产品仍正常运行。

### 涉及模块

- 后端 AI Service (`ai_service.go`) — 全面重写
- Coffee Log 创建/更新逻辑 (`coffee_log_service.go`) — 联动升级
- AI Handler (`ai_handler.go`) — 新增 insight 端点
- 路由 (`router.go`) — 注册 `/ai/flavor-insight` 路由

### 验收标准

- ✅ AI 文案不只描述酸甜苦等参数，具备画面感和 editorial 风格。
- ✅ AI 文案具备生活方式、画面感和 editorial 风格（DeepSeek API 生成）。
- ✅ 同一条记录可以支持不同场景的文案展示（风味摘要 / 生活方式 / 偏好洞察）。
- ✅ DeepSeek API 不可用时自动 fallback 到 Mock 模板文案，产品不中断。
- ✅ 新增 `/ai/flavor-insight` 端点，为统计页和月报提供偏好洞察文案。

---

## Step 6：分享卡片模板优化 ✅ 已完成

### 优先级

P0

### 优化目标

把分享卡片打造成产品传播核心，让用户愿意主动分享到朋友圈、小红书、Instagram 等平台。

### 核心改动

- 增加 Minimal 模板。
- 增加 Magazine 模板。
- 增加 Cafe Receipt 模板。
- 后续可增加 Dark Roast 模板。
- 支持模板切换预览。
- 支持高清 PNG 导出。
- 导出图片带 My Coffee Log 品牌水印。

### 分享卡片字段

- 咖啡照片
- 咖啡名称
- 日期
- 风味标签
- 风味雷达图
- AI 文案
- My Coffee Log 小水印

### 实现详情

- **模板切换器**：新增顶部模板切换栏（极简 / 杂志 / 小票），使用与比例切换器一致的 pill-button 风格，切换时卡片平滑过渡。
- **Minimal 极简模板**：暖奶油色底（#FFF2DB），圆角 4px；仅展示照片 + 咖啡名 + 店名 + AI 文案片段 + 日期/类型；无雷达图，信息极简；品牌水印以淡色 `MY COFFEE LOG / MINIMAL` 呈现。
- **Magazine 杂志模板**：保留原有 double-border 经典设计；照片带类型徽章（espresso/85 底色）；包含风味雷达图；品牌水印 `MY COFFEE LOG / CHRONICLE OF FLAVOR`；底部展示日期 + 心情 emoji。
- **Cafe Receipt 小票模板**：米白纸色底（#FFFEF7），极小圆角 2px + 细边框模拟热敏纸；虚线分隔（border-dashed）；咖啡信息以 receipt line-item 格式展示（COFFEE / TYPE / SHOP / MOOD）；风味分数以 3×2 mono grid 展示；底部装饰性条形码图案；品牌水印 `MY COFFEE LOG` + 条形码装饰。
- **品牌水印**：所有模板均包含 `MY COFFEE LOG` 品牌标识，位置和风格随模板调性变化。
- **导出逻辑升级**：导出文件名包含模板名（`MCL-{template}-{id}-{ratio}.png`）；小票模板使用 `#FFFEF7` 背景色导出，其他模板使用 `#FFF2DB`；底部提示文案动态显示当前模板风格 + 比例 + 渲染倍率。
- **moodEmoji 工具函数**：新增心情到 emoji 的映射（Calm→😌 / Energetic→⚡ / Reflective→💭 / Tired→🥱）。

### 涉及模块

- 分享卡片页 (`ShareCard.vue`) — 全面重写

### 验收标准

- ✅ 至少支持 3 套分享模板（Minimal / Magazine / Cafe Receipt）。
- ✅ 用户可以切换模板预览（pill-button 切换栏）。
- ✅ 用户可以导出高清 PNG（3x 超采样渲染）。
- ✅ 导出的图片包含品牌标识（MY COFFEE LOG）。

---

## Step 7：咖啡心情日记 ✅

### 优先级

P1

### 优化目标

让记录从“咖啡本身”扩展到“喝咖啡时的生活状态”，强化生活方式产品属性。

### 核心改动

- 扩展 mood 字段表达能力。
- 增加心情标签（mood_tags）。
- 增加场景标签（scene_tags）。
- 增加搭配标签（pairing_tags）。
- 在详情页展示生活方式标签。
- 后续支持按心情或场景统计。

### 标签示例

- 心情：Calm、Focused、Tired、Happy、Rainy、Slow、Productive
- 场景：Morning、Office、Weekend、Cafe、Travel、Home、Study
- 搭配：Book、Music、Work、Dessert、Alone、Friends

### 涉及模块

- 创建页
- 咖啡详情页
- Coffee Log 数据模型
- Coffee Log API
- 时间线筛选
- 统计页

### 验收标准

- 用户可以为一条记录选择心情、场景、搭配。
- 详情页能展示这些生活方式标签。
- 后续统计页可以基于这些标签生成洞察。

### 实现记录

- **后端**：CoffeeLog 模型新增 `mood_tags`、`scene_tags`、`pairing_tags` 三个 JSON 字段（varchar(500)），通过 GORM AutoMigrate 自动迁移。Create/Update 请求结构体新增对应字段，后端校验标签值在预设白名单内且单类最多 5 个。
- **前端**：`constants/coffee.ts` 新增 `LIFESTYLE_MOOD_TAGS`、`LIFESTYLE_SCENE_TAGS`、`LIFESTYLE_PAIRING_TAGS` 常量。API 层 `CoffeeLogDTO` 和 `CreateCoffeeLogParams` 新增对应字段，`toCoffeeLog` 自动 JSON.parse。Store 接口同步更新。
- **创建页**：快速记录和精细记录（步骤 3）均新增生活标签多选区域，心情/场景/搭配分别以 amber/sky/rose 色系区分。
- **详情页**：新增「生活标签 / Lifestyle」区块，按分类展示已选标签，无标签时自动隐藏。

---

## Step 8：统计页升级为个人咖啡画像 ✅

### 优先级

P1

### 优化目标

让统计页从数据展示升级为“个人咖啡画像”，提升情绪价值和分享价值。

### 核心改动

- 新增 Your Coffee Personality 模块。
- 根据历史记录生成咖啡人格标签。
- 为每个人格标签生成解释文案。
- 支持将咖啡人格生成分享图。

### 人格示例

- Citrus Minimalist：偏爱清爽、明亮、酸质突出的咖啡。
- Creamy Comfort Seeker：偏爱奶咖、焦糖、坚果、顺滑口感。
- Slow Morning Brewer：常在早晨记录手冲，偏好安静、轻盈的风味。
- Urban Latte Lover：偏好咖啡店、拿铁、通勤场景。
- Bold Espresso Purist：追求浓烈、醇厚、苦甜交织的深度体验。
- Rainy Day Reader：阴雨天、独处、阅读的安静咖啡时光。
- Social Weekend Explorer：周末、朋友、甜点的社交咖啡。
- Productive Hustler：专注、高效、工作的燃料咖啡。
- Coffee Explorer：正在探索咖啡世界的入门者（兜底人格）。

### 涉及模块

- 统计页
- Stats API
- Coffee Log Store
- AI 偏好洞察
- 分享卡片

### 验收标准

- 统计页展示用户咖啡人格标签。
- 人格标签有对应说明文案。
- 用户可以将咖啡人格生成分享图。

### 实现记录

- **后端**：`stats_repository.go` 新增 `GetLifestyleTagCounts` 方法，聚合 mood_tags/scene_tags/pairing_tags JSON 字段中的标签使用频次。`stats_service.go` 新增 `GetPersonality` 方法，基于 8 条人格规则（匹配 + 评分）从风味特征、咖啡类型偏好、生活方式标签中推导出最多 3 个人格标签，无匹配时兜底为 Coffee Explorer。`stats_handler.go` 新增 `GetPersonality` 接口，路由 `GET /api/v1/stats/personality`。
- **前端**：`api/stats.ts` 新增 `PersonalityTag`、`PersonalityResponse` 类型和 `getPersonality` API 调用。Store 新增 `personalities` ref 和 `fetchPersonality` action。`Stats.vue` 新增「你的咖啡人格 / Coffee Personality」模块，包含主人格卡片（图标 + 标题 + 副标题 + 描述）、次人格标签列表、分享按钮（Web Share API / 剪贴板兜底）。

---

## Step 9：复刻上一杯 ✅

### 优先级

P1

### 优化目标

提高重复记录效率，让用户可以快速复用历史记录。

### 核心改动

- 在详情页增加 Brew Again 入口。
- 在时间线卡片增加 Copy This Log 入口。
- 点击后进入创建页并自动填充原记录信息。
- 保存时生成新记录，不覆盖原记录。

### 自动带入字段

- 咖啡类型
- 风味参数
- 风味标签
- 店铺
- 备注结构

### 涉及模块

- 咖啡详情页
- 时间线
- 创建页
- Coffee Log Store
- 路由参数

### 验收标准

- 用户可以从历史记录发起复刻。
- 创建页能自动填充原记录信息。
- 保存后生成一条新的咖啡记录。

### 实现记录

- **前端**：`CoffeeDetail.vue` 将"再记一杯"改为"复刻这杯"按钮，链接到 `/create?from_log_id={id}`。`Timeline.vue` 三种卡片类型均新增 hover 显示的复制按钮（Copy icon），链接到 `/create?from_log_id={id}`。`CreateCoffeeLog.vue` 新增 `onMounted` 钩子，读取 `route.query.from_log_id`，从 store 或 API 获取源记录，自动填充 quickForm 和 form 的咖啡类型、风味参数、风味标签、心情、店铺、备注、生活方式标签等字段。保存时生成新记录，不覆盖原记录。

---

## Step 10：月度咖啡回顾 ✅ 已完成

### 优先级

P1

### 优化目标

通过月度总结提升用户留存，并创造可分享内容。

### 核心改动

- 增加月度咖啡回顾模块。
- 聚合本月记录数据。
- 生成本月咖啡关键词。
- 生成 AI 月度总结文案。
- 支持导出月度回顾卡片。

### 月报内容

- 本月喝了几杯
- 最常喝的类型
- 最常出现的风味
- 最常喝咖啡的日期或时段
- 本月咖啡关键词
- 月度总结卡片

### 涉及模块

- 统计页
- Stats API
- AI 文案
- 分享卡片
- 月度数据聚合

### 验收标准

- 用户可以查看本月咖啡回顾。
- 月报包含数据总结和 AI 文案。
- 月报可以导出为分享图片。

### 实现记录

- **后端**：`stats_repository.go` 新增 `GetMonthCountByMonth`、`GetFavoriteCoffeeTypeByMonth`、`GetTopCoffeeTypesByMonth`、`GetTopFlavorTagsByMonth`、`GetTopCoffeeNamesByMonth`、`GetTopDrinkWeekdayByMonth`、`GetLifestyleTagCountsByMonth`、`GetFlavorProfileByMonth` 8 个按月聚合查询方法，支持传入 `YYYY-MM` 格式月份参数。新增 `CoffeeTypeCount`、`CoffeeNameCount`、`WeekdayCount` 数据结构。
- **后端**：`stats_service.go` 新增 `GetMonthlyReview` 方法，聚合月度杯数、偏好类型、类型分布、风味标签 Top N、常喝咖啡名 Top N、最常喝咖啡的星期、心情/场景/搭配标签、月度风味雷达数据。新增 `generateMonthlyKeywords` 方法，根据偏好类型、风味、心情、场景、星期自动生成 6 个中文关键词（如"拿铁爱好者""柑橘风味""平静时光"等）。新增 `MonthlyReviewResponse` 结构体包含完整月报数据。
- **后端**：`ai_service.go` 新增 `GenerateMonthlyReview` 方法和 `GenerateMonthlyReviewForUser` 便捷方法，支持 DeepSeek API 生成 editorial 风格月度回顾文案，API 不可用时自动 fallback 到 `generateMockMonthlyReview` 模板文案。System Prompt 要求 3-5 句话、80-150 字、画面感和比喻风格。新增 `GET /api/v1/ai/monthly-review` 路由。
- **后端**：`stats_handler.go` 新增 `GetMonthlyReview` handler，接受 `month` 查询参数（默认当前月）。新增 `GET /api/v1/stats/monthly-review` 路由。
- **前端**：`api/stats.ts` 新增 `MonthlyReviewData`、`MonthlyReviewFlavorTag`、`MonthlyReviewCoffeeType`、`MonthlyReviewCoffeeName`、`MonthlyReviewLifestyleTag`、`MonthlyReviewFlavorProfile`、`MonthlyReviewAIResponse` 类型和 `getMonthlyReview`、`getMonthlyReviewAI` API 调用。
- **前端**：Store 新增 `monthlyReview` ref、`monthlyReviewAI` ref 和 `fetchMonthlyReview` action，并行请求月报数据和 AI 文案。
- **前端**：新增 `MonthlyReview.vue` 页面，包含月度杯数大数字 + 关键词卡片、AI Editorial 月度回顾文案（装饰性大引号 + closing ornament）、冲煮偏好分布条形图、风味图谱标签云、月度风味雷达图、常喝咖啡 Top 排行、生活标签（心情/场景/搭配三色系）、最常喝咖啡的星期、分享按钮（Web Share API / 剪贴板兜底）。支持左右箭头切换月份，空月展示引导状态。
- **前端**：`Stats.vue` 新增「月度咖啡回顾」入口卡片（📅 图标 + editorial 渐变背景 + ChevronRight 箭头），链接到 `/monthly-review`。

---

## Step 11：咖啡店收藏 ✅ 已完成

### 优先级

P1

### 优化目标

沉淀用户长期咖啡消费场景，让产品从“记录咖啡”扩展到“记录咖啡地点”。

### 核心改动

- 新增咖啡店数据维度。
- 创建记录时支持复用历史店铺。
- 增加咖啡店列表页。
- 增加咖啡店详情页。
- 统计到访次数和关联记录。

### 店铺字段

- 店铺名称
- 店铺地址
- 店铺评分
- 到访次数
- 店铺照片
- 最近一次到访

### 涉及模块

- Coffee Shop 数据模型
- 创建页店铺选择
- 咖啡店列表页
- 咖啡店详情页
- 统计页

### 验收标准

- 用户可以复用历史店铺名称。
- 用户可以查看自己记录过的咖啡店。
- 店铺详情能展示到访次数和关联咖啡记录。

### 实现记录

- **后端**：新增 `model/coffee_shop.go` CoffeeShop 数据模型，包含 Name、Address、Rating、ImageURL、VisitCount、LastVisitAt 等字段，支持软删除。
- **后端**：新增 `repository/coffee_shop_repository.go`，包含 Create、FindByID、FindList（分页+搜索）、FindByName、Update、Delete、FindShopNames、IncrementVisitCount 等方法。新增 `repository/coffee_log_repository.go` 的 `FindByShopName` 方法按店铺名称查询关联咖啡记录。
- **后端**：新增 `service/coffee_shop_service.go`，包含 CRUD 操作、`EnsureShopForLog`（创建咖啡记录时自动创建或更新店铺记录并递增到访次数）、`GetRelatedLogs`（获取店铺关联咖啡记录）等业务逻辑。
- **后端**：`service/coffee_log_service.go` 的 Create 方法集成 `EnsureShopForLog`，用户创建咖啡记录时自动维护店铺收藏数据。
- **后端**：新增 `handler/coffee_shop_handler.go`，包含 Create、GetByID、GetList、Update、Delete、GetShopNames、GetRelatedLogs 七个 handler。新增路由组 `coffee-shops`，包含 `POST /`、`GET /`、`GET /names`、`GET /:id`、`PUT /:id`、`DELETE /:id`、`GET /:id/logs`。
- **后端**：`main.go` 新增 CoffeeShop AutoMigrate、shopRepo、shopService、shopHandler 依赖注入，调整服务初始化顺序避免循环依赖。
- **前端**：新增 `api/coffeeShop.ts`，定义 CoffeeShop、CreateCoffeeShopParams、UpdateCoffeeShopParams 类型和 getCoffeeShops、getCoffeeShopById、createCoffeeShop、updateCoffeeShop、deleteCoffeeShop、getShopNames、getShopLogs API 调用。
- **前端**：新增 `CoffeeShops.vue` 咖啡店列表页，包含搜索栏、店铺卡片（名称/地址/评分/到访次数/最近到访）、添加店铺弹窗（名称/地址/星级评分）、空状态引导。
- **前端**：新增 `CoffeeShopDetail.vue` 咖啡店详情页，包含封面图/店铺信息卡（名称/地址/评分/到访统计/最近到访/关联记录数）、关联咖啡记录列表、编辑店铺弹窗（名称/地址/评分/删除）。
- **前端**：`router/index.ts` 新增 `/coffee-shops` 和 `/coffee-shops/:id` 路由。
- **前端**：`Profile.vue` 系统偏好菜单新增「咖啡店收藏」入口，链接到 `/coffee-shops`。
- **前端**：`CreateCoffeeLog.vue` 精细记录模式 Step 3 店铺输入框增加自动补全下拉，聚焦时加载历史店铺名称列表，输入时实时过滤，点击补全项自动填充。

---

## Step 12：咖啡豆档案 ✅ 已完成

### 优先级

P2

### 优化目标

服务更专业的手冲用户，但不影响普通用户的轻量记录体验。

### 核心改动

- 在精细记录中增加咖啡豆字段。
- 支持保存和复用咖啡豆档案。
- 详情页展示豆子信息。
- 后续统计用户偏好的产地、处理法、烘焙度。

### 字段建议

- 产地
- 处理法
- 烘焙度
- 烘焙商
- 豆子名称
- 冲煮方式
- 粉水比
- 水温
- 研磨度

### 涉及模块

- 创建页高级字段
- Coffee Bean 数据模型
- Coffee Log 数据模型
- 详情页
- 统计页

### 验收标准

- 普通用户不填写豆子信息也能完成记录。
- 高级用户可以记录并复用咖啡豆信息。
- 咖啡详情页能展示豆子档案。

### 实现记录

- **后端**：新增 `model/coffee_bean.go` CoffeeBean 数据模型，包含 Name、Origin、ProcessingMethod、RoastLevel、Roaster、ImageURL、UsageCount 等字段，支持软删除。
- **后端**：`model/coffee_log.go` 新增 BeanID（外键关联 CoffeeBean）、BrewRatio、WaterTemp、GrindSize 字段，新增 Bean 关联预加载。
- **后端**：新增 `repository/coffee_bean_repository.go`，包含 Create、FindByID、FindList（分页+搜索）、FindByName、Update、Delete、FindBeanNames、IncrementUsageCount 等方法。
- **后端**：新增 `service/coffee_bean_service.go`，包含 CRUD 操作、`EnsureBeanForLog`（创建咖啡记录时自动创建或查找咖啡豆记录并递增使用次数）等业务逻辑。
- **后端**：`service/coffee_log_service.go` 新增 beanService 依赖，Create 方法集成 `EnsureBeanForLog`，Create/Update 请求结构体新增 BeanID、BeanName、BrewRatio、WaterTemp、GrindSize 字段。
- **后端**：`repository/coffee_log_repository.go` FindByID 新增 Preload("Bean") 预加载。
- **后端**：新增 `handler/coffee_bean_handler.go`，包含 Create、GetByID、GetList、Update、Delete、GetBeanList 六个 handler。新增路由组 `coffee-beans`，包含 `POST /`、`GET /`、`GET /list`、`GET /:id`、`PUT /:id`、`DELETE /:id`。
- **后端**：`main.go` 新增 CoffeeBean AutoMigrate、beanRepo、beanService、beanHandler 依赖注入。
- **前端**：新增 `api/coffeeBean.ts`，定义 CoffeeBean、CreateCoffeeBeanParams、UpdateCoffeeBeanParams 类型和 getCoffeeBeans、getCoffeeBeanById、createCoffeeBean、updateCoffeeBean、deleteCoffeeBean、getBeanList API 调用。
- **前端**：`api/coffeeLog.ts` CoffeeLogDTO 新增 bean_id、bean、brew_ratio、water_temp、grind_size 字段，toCoffeeLog 映射新增对应字段，CreateCoffeeLogParams 和 toCreatePayload 新增 bean_id、bean_name、brew_ratio、water_temp、grind_size。
- **前端**：`stores/coffeeLog.ts` CoffeeLog 接口新增 bean_id、bean（CoffeeBeanInfo）、brew_ratio、water_temp、grind_size 字段，NewCoffeeLog 接口新增 bean_id、bean_name、brew_ratio、water_temp、grind_size 可选字段。
- **前端**：`CreateCoffeeLog.vue` 精细记录模式 Step 2 新增冲煮参数区域（粉水比/水温/研磨度输入框），Step 3 新增咖啡豆档案区域（已保存豆子下拉选择 + 手动填写豆子名称/产地/处理法/烘焙度/烘焙商），提交时自动创建内联豆子记录。
- **前端**：`CoffeeDetail.vue` 新增咖啡豆档案展示区域（豆子名称/产地/处理法/烘焙度/烘焙商 + 冲煮参数粉水比/水温/研磨度），仅在有豆子信息时显示。

---

## Step 13：冲煮参数记录

### 优先级

P2

### 优化目标

为专业用户提供更完整的冲煮记录能力。

### 核心改动

- 在高级模式中增加冲煮参数。
- 参数区域默认折叠。
- 详情页以清晰排版展示专业参数。

### 参数建议

- 水温
- 粉量
- 水量
- 粉水比
- 研磨度
- 萃取时间
- 冲煮器具
- 滤纸类型

### 涉及模块

- 创建页高级模式
- Coffee Log 数据模型
- 后端接口
- 详情页

### 验收标准

- 冲煮参数默认折叠，不影响快速记录。
- 用户可以保存完整冲煮参数。
- 详情页能以清晰排版展示专业参数。

---

## Step 14：咖啡地图

### 优先级

P2

### 优化目标

基于咖啡店与位置数据，扩展城市咖啡生活记录场景。

### 核心改动

- 基于咖啡店位置生成地图点位。
- 展示用户咖啡足迹。
- 支持从地图进入店铺详情。
- 支持按城市或区域筛选。
- 无位置数据时展示空状态。

### 涉及模块

- 咖啡店数据
- 位置字段
- 地图组件
- 店铺详情页

### 验收标准

- 用户可以通过地图查看自己的咖啡足迹。
- 地图点位可以进入对应店铺详情。
- 无位置数据时有合理空状态。

---

## Step 15：真实 AI 接入 ✅ 已完成

### 优先级

P2

### 优化目标

从 Mock AI 升级为真实模型能力，生成更个性化的总结、月报与分享文案。

### 核心改动

- 后端 AI Service 支持真实模型调用。
- 支持环境变量配置。
- 保留 Mock AI 作为降级方案。
- AI 请求失败不影响用户保存记录。

### AI 能力范围

- 单杯咖啡风味摘要
- 分享卡片文案
- 月度咖啡回顾
- 个人咖啡画像
- 偏好变化洞察

### 配置项建议

- `OPENAI_API_KEY`
- `OPENAI_BASE_URL`
- `OPENAI_MODEL`

### 涉及模块

- 后端 AI Service
- 配置管理
- Coffee Log Service
- Stats Service
- 错误降级逻辑

### 验收标准

- 未配置 API Key 时可以继续使用 Mock AI。
- 配置 API Key 后可以生成真实 AI 文案。
- AI 请求失败时不会影响用户保存记录。

### 实现记录

- **后端**：`service/ai_service.go` 已有 `callDeepSeekAPI()` 函数支持真实 OpenAI 兼容 API 调用，新增 `ExternalAIEnabled()` 公共方法供 handler 调用。
- **后端**：`service/ai_service.go` 新增 `GenerateShareCopy` 分享卡片文案生成（含 mock 降级）、`GenerateCoffeeProfileForUser` 个人咖啡画像 AI 生成、`GeneratePreferenceInsightForUser` 偏好变化洞察 AI 生成，三个新 AI 能力均支持真实 API 调用 + Mock 降级。
- **后端**：`handler/ai_handler.go` 新增 `GetAIStatus`（返回 AI 启用状态和模型名）、`GenerateShareCopy`、`GenerateCoffeeProfile`、`GeneratePreferenceInsight` 四个 handler。
- **后端**：`router/router.go` AI 路由组新增 `GET /status`、`POST /share-copy`、`POST /coffee-profile`、`POST /preference-insight` 四个端点。
- **后端**：`.env.example` 新增 `AI_ENABLED` 和 `OPENAI_REQUEST_TIMEOUT_SECONDS` 配置项说明。
- **前端**：`api/stats.ts` 新增 `AIStatus`、`ShareCopyRequest/Response`、`CoffeeProfileResponse`、`PreferenceInsightResponse` 类型和 `getAIStatus`、`generateShareCopy`、`generateCoffeeProfile`、`generatePreferenceInsight` API 调用。
- **前端**：`stores/coffeeLog.ts` 新增 `aiStatus`、`coffeeProfileAI`、`preferenceInsight` 状态和 `fetchAIStatus`、`fetchCoffeeProfile`、`fetchPreferenceInsight` actions。
- **前端**：`Stats.vue` 新增 AI 咖啡画像展示区、偏好洞察展示区、AI 引擎状态指示器（显示启用/禁用状态和模型名），onMounted 自动获取 AI 状态和文案。
- **前端**：`ShareCard.vue` 新增「AI 生成分享文案」按钮，调用后端 AI 生成分享短文案，支持复制到剪贴板。

---

## Step 16：轻量成就系统

### 优先级

P2

### 优化目标

用克制的方式增强长期记录动力，不破坏产品高级感。

### 核心改动

- 增加轻量成就规则。
- 成就入口放在个人中心或统计页。
- 达成后低打扰提示。
- 成就视觉保持低饱和、高级、克制。

### 成就示例

- 连续记录 3 天
- 本月记录 10 杯
- 解锁 5 种风味
- 第一次记录冷萃
- 第一次生成分享卡片

### 涉及模块

- 个人中心
- 统计页
- 成就规则
- Coffee Log 数据统计

### 验收标准

- 用户达成条件后可以看到成就。
- 成就提示不打断主流程。
- 成就视觉风格与整体产品调性一致。

---

## 2. 推荐执行顺序

### 第一阶段：P0 核心闭环优化

1. Quick Log 快速记录模式
2. 首页咖啡生活仪表盘
3. 记录完成后的即时正反馈
4. 咖啡详情页视觉强化
5. AI 文案风格升级
6. 分享卡片模板优化

### 第二阶段：P1 留存与复访优化

1. 咖啡心情日记
2. 统计页升级为个人咖啡画像
3. 复刻上一杯
4. 月度咖啡回顾
5. 咖啡店收藏

### 第三阶段：P2 深度能力扩展

1. 咖啡豆档案
2. 冲煮参数记录
3. 咖啡地图
4. 真实 AI 接入
5. 轻量成就系统

---

## 3. 不建议优先做的方向

### 3.1 不要一开始做得过于专业

粉水比、水温、萃取时间等字段对普通用户有门槛，可以做，但应该放在高级模式中，避免影响主流程。

### 3.2 不要把首页做成普通管理后台

这个项目的核心竞争力是审美和情绪价值，不是表格效率。首页应该更像生活方式仪表盘，而不是数据管理页面。

### 3.3 不要过早做社交关系链

好友、关注、点赞、评论会显著增加复杂度。前期更适合做好外部分享，而不是自建社区。

### 3.4 不要让 AI 显得太工具化

AI 应该像一个高级生活方式编辑，而不是参数分析器。AI 输出需要服务于审美、情绪、分享和个人洞察。

---

## 4. 最终总结

My Coffee Log 的后续优化建议应按照“一个优化点一个步骤”的方式推进。

第一阶段优先完成 Quick Log、首页仪表盘、记录后正反馈、详情页强化、AI 文案升级和分享卡片模板。

第二阶段重点提升留存和复访，包括咖啡心情日记、个人咖啡画像、复刻上一杯、月度咖啡回顾和咖啡店收藏。

第三阶段再扩展咖啡豆档案、冲煮参数、咖啡地图、真实 AI 接入和轻量成就系统。