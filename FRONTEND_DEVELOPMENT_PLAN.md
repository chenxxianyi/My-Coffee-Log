# My Coffee Log 前端开发方案

## 1. 前端目标

My Coffee Log 前端需要呈现一款 **Nordic Minimal + Editorial** 风格的咖啡生活记录 App。

它不是普通 Dashboard，也不是传统 CRUD 管理系统，而是一个以图片、留白、排版和情绪记录为核心的生活方式产品。

核心体验：

> 用户用 3 秒记录一杯咖啡，查看像杂志一样的咖啡时间线，并生成高级分享卡片。

---

## 2. 前端技术栈

- Vue 3
- Vite
- TypeScript
- Pinia
- Vue Router
- Tailwind CSS
- SVG 自定义雷达图或 ECharts
- html2canvas
- Capacitor

### 推荐实现选择

- **雷达图**：优先使用 SVG 自定义组件，方便控制杂志感视觉。
- **分享卡片导出**：使用 Vue 组件渲染卡片，再通过 html2canvas 导出 PNG。
- **移动端适配**：先完成 Web MVP，再接入 Capacitor 打包移动端。

---

## 3. 前端目录结构

```text
frontend/
  src/
    api/
      auth.ts
      coffeeLog.ts
      stats.ts
      ai.ts
    assets/
    components/
      base/
      coffee/
      charts/
      share/
    layouts/
    pages/
      Home.vue
      Login.vue
      Register.vue
      CreateCoffeeLog.vue
      CoffeeDetail.vue
      Timeline.vue
      Stats.vue
      Profile.vue
    router/
    stores/
      auth.ts
      coffeeLog.ts
    styles/
    utils/
    App.vue
    main.ts
```

---

## 4. 前端模块职责

### 4.1 api

负责封装后端 API 请求。

建议文件：

- `auth.ts`：注册、登录、获取当前用户、更新用户信息。
- `coffeeLog.ts`：咖啡记录创建、列表、详情、更新、删除。
- `stats.ts`：统计数据接口。
- `ai.ts`：AI 风味总结接口。

统一处理：

- Base URL
- Authorization Token
- 请求错误
- 统一响应格式

### 4.2 stores

使用 Pinia 管理全局状态。

建议 Store：

- `auth.ts`：token、当前用户、登录状态、登录、登出、拉取用户信息。
- `coffeeLog.ts`：咖啡记录列表、当前详情、筛选条件、创建与刷新逻辑。

### 4.3 components/base

基础 UI 组件。

建议组件：

- `BaseButton.vue`
- `BaseInput.vue`
- `BaseTextarea.vue`
- `BaseTag.vue`
- `BaseModal.vue`
- `BaseEmpty.vue`
- `BaseLoading.vue`

设计原则：

- 少阴影
- 细边框
- 大留白
- 克制交互
- 温暖低饱和色彩

### 4.4 components/coffee

咖啡业务组件。

建议组件：

- `CoffeeTypeSelector.vue`
- `CoffeeLogCard.vue`
- `CoffeeLogHero.vue`
- `FlavorSliderGroup.vue`
- `FlavorTagSelector.vue`
- `MoodSelector.vue`
- `TimelineMonthGroup.vue`

### 4.5 components/charts

图表组件。

建议组件：

- `FlavorRadarChart.vue`

要求：

- 使用六个维度：酸度、苦感、甜感、醇厚度、香气、余韵。
- 0-5 分映射到雷达半径。
- 网格线柔和。
- 色彩低饱和。
- 不做商业 BI 或电竞风格。

### 4.6 components/share

分享卡片相关组件。

建议组件：

- `ShareCardPreview.vue`
- `CoffeeShareCard.vue`

功能：

- 选择卡片比例。
- 渲染卡片预览。
- 调用 html2canvas 导出图片。
- 支持本地下载。

---

## 5. 页面开发方案

## 5.1 Login.vue

### 页面目标

登录页应是品牌入口，而不是普通表单页。

### 内容

- 产品名称：My Coffee Log
- 品牌文案
- 邮箱输入框
- 密码输入框
- 登录按钮
- 注册入口

### UI 要求

- 极简
- 高级
- 温暖纯色或咖啡摄影背景
- 大标题排版
- 少装饰
- 不使用普通后台登录页风格

### 交互

- 登录成功后保存 token。
- 拉取当前用户信息。
- 跳转 Home。

---

## 5.2 Register.vue

### 页面目标

让用户快速创建账号，保持和 Login 一致的品牌感。

### 内容

- 邮箱
- 密码
- 昵称
- 注册按钮
- 登录入口

### 交互

- 注册成功后可直接登录，或跳转登录页。
- 参数错误时展示轻量错误提示。

---

## 5.3 Home.vue

### 页面目标

首页应该像高级生活方式杂志封面，而不是普通 Dashboard。

### 展示内容

- 今日咖啡入口
- 快速创建按钮
- 最近一条咖啡记录
- 最近 7 天咖啡记录
- AI 今日风味摘要
- 本月咖啡次数

### 布局建议

- 顶部使用大标题和日期。
- 最近记录使用大图排版。
- 本月次数用轻量文字数字展示，不做厚重统计卡片。
- 快速创建按钮应明显但克制。

### 数据来源

- `GET /api/v1/coffee-logs`
- `GET /api/v1/stats/overview`

---

## 5.4 CreateCoffeeLog.vue

### 页面目标

实现 3 秒可完成最小记录的三步式创建流程。

### 第一步：基础信息

- 选择咖啡类型
- 上传图片或填写图片 URL
- 填写咖啡名称

最小记录可以只选择类型并提交默认值，其他字段不阻塞用户。

### 第二步：风味强度

通过滑动条选择：

- 酸度 `acidity`
- 苦感 `bitterness`
- 甜感 `sweetness`
- 醇厚度 `body`
- 香气 `aroma`
- 余韵 `aftertaste`

范围：0-5。

### 第三步：情绪和标签

- 选择风味标签
- 选择心情
- 填写店铺
- 填写备注

### 提交流程

- 组装 Coffee Log 请求体。
- 调用 `POST /api/v1/coffee-logs`。
- 后端自动生成 AI Summary。
- 创建成功后跳转 CoffeeDetail。

---

## 5.5 CoffeeDetail.vue

### 页面目标

详情页应像生活方式杂志内页。

### 展示内容

- 咖啡大图
- 咖啡名称
- 日期
- 店铺
- 地点
- 风味雷达图
- 风味标签
- AI 风味描述
- 心情
- 备注
- 编辑按钮
- 删除按钮
- 生成分享卡片按钮

### 数据来源

- `GET /api/v1/coffee-logs/:id`

### 交互

- 点击分享按钮打开 `ShareCardPreview`。
- 点击编辑进入编辑模式或复用创建页。
- 删除前二次确认。

---

## 5.6 Timeline.vue

### 页面目标

Timeline 是用户的咖啡生活日记，不是普通列表。

### 展示内容

- 按月份分组的咖啡记录
- 图片优先的记录卡片
- 日期、咖啡名、店铺、标签

### 筛选能力

- 按月份筛选
- 按咖啡类型筛选
- 按风味标签筛选

### UI 要求

- 大量留白
- 图片优先
- 时间流感
- 像一本咖啡日记
- 避免表格和后台列表风格

---

## 5.7 Stats.vue

### 页面目标

展示用户的咖啡偏好，但保持杂志感，不做商业 BI 页面。

### 展示内容

- 本月喝咖啡次数
- 总记录数
- 最常喝的咖啡类型
- 最常出现的风味标签
- 风味偏好平均值
- 月度记录趋势

### 数据来源

- `GET /api/v1/stats/overview`
- `GET /api/v1/stats/flavor-profile`
- `GET /api/v1/stats/monthly`

### UI 要求

- 图表柔和
- 少用强边框和重色块
- 数字展示更像编辑排版
- 保持温暖、安静、高级

---

## 5.8 Profile.vue

### 页面目标

用户资料设置。

### 功能

- 展示当前用户邮箱
- 修改昵称
- 修改头像 URL
- 退出登录

### 数据来源

- `GET /api/v1/users/me`
- `PUT /api/v1/users/me`

---

## 6. 路由设计

建议路由：

```text
/login
/register
/
/create
/coffee-logs/:id
/timeline
/stats
/profile
```

### 路由守卫

- `/login` 和 `/register` 不需要登录。
- 其他页面全部需要登录。
- 未登录访问受保护页面时跳转 `/login`。
- 已登录访问 `/login` 或 `/register` 时可跳转首页。

---

## 7. API 对接方案

### 7.1 统一响应格式

成功：

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

失败：

```json
{
  "code": 40001,
  "message": "error message",
  "data": null
}
```

### 7.2 Token 处理

前端登录成功后保存 token。

后续请求统一添加：

```http
Authorization: Bearer <token>
```

### 7.3 分页处理

分页返回格式：

```json
{
  "list": [],
  "pagination": {
    "page": 1,
    "page_size": 10,
    "total": 100
  }
}
```

Timeline 和列表页面需要基于该格式做分页加载或筛选刷新。

---

## 8. Coffee Log 前端数据结构

```ts
interface CoffeeLog {
  id: number
  coffee_name: string
  coffee_type: string
  shop_name?: string
  location?: string
  image_url?: string
  drink_date: string
  mood?: string
  notes?: string
  acidity: number
  bitterness: number
  sweetness: number
  body: number
  aroma: number
  aftertaste: number
  ai_summary?: string
  flavor_tags?: FlavorTag[]
  created_at: string
  updated_at: string
}
```

```ts
interface FlavorTag {
  id: number
  name: string
  label: string
  color: string
}
```

---

## 9. 设计系统

### 9.1 视觉风格

整体采用：

- Nordic Minimal
- Editorial
- Premium
- Warm
- Calm
- Coffee Lifestyle
- Typography Driven

### 9.2 配色

| 名称 | 色值 |
| --- | --- |
| Warm White | `#F7F3EC` |
| Cream | `#EFE7DA` |
| Latte Beige | `#D7C4A8` |
| Coffee Brown | `#7A5638` |
| Deep Espresso | `#2A1E17` |
| Charcoal | `#1E1E1E` |
| Soft Gray | `#A8A29A` |

### 9.3 布局原则

- 大量留白
- 大图优先
- 非对称排版
- 杂志式层级
- 少用传统卡片网格
- 少用重阴影
- 不做普通 Dashboard
- 不做普通 SaaS 风
- 不做 Material Design 风
- 不使用廉价渐变

---

## 10. 风味雷达图方案

### 数据维度

- acidity 酸度
- bitterness 苦感
- sweetness 甜感
- body 醇厚度
- aroma 香气
- aftertaste 余韵

### 实现建议

使用 SVG 自定义雷达图组件。

### 视觉规则

- 半径按 0-5 映射。
- 网格线使用低透明度深咖啡色或灰色。
- 填充色使用 Latte Beige 或 Coffee Brown 的低透明度。
- 标签字体小而精致。
- 不使用高饱和荧光色。

---

## 11. 分享卡片方案

### 11.1 功能目标

用户可以将一条 Coffee Log 生成精美图片并保存到本地。

### 11.2 卡片内容

- 咖啡图片
- 咖啡名称
- 日期
- 店铺
- 风味雷达图
- 风味标签
- AI 文案
- My Coffee Log 品牌标识

### 11.3 支持尺寸

- 1:1 小红书 / Instagram
- 3:4 小红书封面
- 9:16 Story

建议导出尺寸：

- 1:1：1080 × 1080
- 3:4：1080 × 1440
- 9:16：1080 × 1920

### 11.4 导出流程

1. CoffeeDetail 点击生成分享卡片。
2. 打开 ShareCardPreview。
3. 用户选择尺寸。
4. 渲染 CoffeeShareCard。
5. 使用 html2canvas 捕获 DOM。
6. 生成 PNG。
7. 下载到本地。

### 11.5 风险

html2canvas 对跨域图片敏感。

MVP 建议：

- 使用同源图片。
- 或先允许用户填写可跨域访问的图片 URL。
- 后续加入后端上传能力。

---

## 12. 前端开发阶段

## 第一阶段：前端初始化与登录注册

### 任务

- 初始化 Vue 3 + Vite + TypeScript。
- 配置 Tailwind CSS。
- 配置 Vue Router。
- 配置 Pinia。
- 封装 API Client。
- 实现 Login 页面。
- 实现 Register 页面。
- 实现 Auth Store。
- 实现路由守卫。
- 实现基础布局。

### 验收

- 用户可以注册。
- 用户可以登录。
- 登录后可以进入首页。
- 刷新页面后可恢复登录状态。

---

## 第二阶段：Coffee Log 创建、详情和 Timeline

### 任务

- 实现 CreateCoffeeLog 页面。
- 实现三步式记录流程。
- 实现风味滑动条。
- 实现风味标签选择。
- 实现 CoffeeDetail 页面。
- 实现 Timeline 页面。
- 实现月份、咖啡类型、标签筛选。
- 实现编辑和删除入口。

### 验收

- 登录用户可以创建咖啡记录。
- 可以查看记录详情。
- 可以查看自己的时间线。
- 可以编辑和删除自己的记录。

---

## 第三阶段：AI、雷达图和统计页

### 任务

- 实现 FlavorRadarChart。
- CoffeeDetail 展示雷达图。
- CoffeeDetail 展示 AI Summary。
- 实现 Stats 页面。
- Home 展示最近记录和本月统计。

### 验收

- 咖啡详情展示风味雷达图。
- 咖啡详情展示 AI 文案。
- 统计页展示本月次数、常喝类型、常见标签和风味偏好。

---

## 第四阶段：分享卡片和 UI 打磨

### 任务

- 实现 ShareCardPreview。
- 实现 CoffeeShareCard。
- 支持 1:1、3:4、9:16 三种比例。
- 使用 html2canvas 导出图片。
- 打磨 Home 杂志封面感。
- 打磨 Timeline 咖啡日记感。
- 打磨 CoffeeDetail 杂志内页感。
- 打磨 Login / Register 品牌入口感。
- 做响应式适配。
- 预留 Capacitor 移动端打包能力。

### 验收

- 可以生成分享卡片图片。
- UI 整体具有高级生活方式杂志风。
- 移动端浏览体验良好。

---

## 13. 前端 MVP 验收标准

1. 用户可以注册和登录。
2. 登录后可以创建咖啡记录。
3. 可以查看咖啡记录列表和详情。
4. 可以编辑、删除自己的咖啡记录。
5. 咖啡详情展示风味雷达图。
6. 咖啡详情展示 AI Mock 风味总结。
7. 可以查看简单统计。
8. 可以生成分享卡片图片。
9. 页面风格符合 Nordic Minimal + Editorial。
10. 前端可以通过 Docker Compose 启动。

---

## 14. 前端重点风险

### 14.1 容易做成普通后台

需要避免：

- 普通表格列表
- 普通 Dashboard 卡片堆叠
- 重阴影
- 高饱和渐变
- Material Design 默认风格

### 14.2 分享卡片导出失败

风险来源：

- 跨域图片
- 字体未加载完成
- DOM 尺寸缩放

建议：

- 导出前等待图片和字体加载。
- 使用固定导出尺寸。
- 图片尽量走同源资源。

### 14.3 三步创建流程过重

建议：

- 所有非必要字段都不阻塞提交。
- 提供默认日期和默认风味值。
- 用户可以跳过部分步骤快速完成。

---

## 15. 前端最终目标

前端最终要让用户感觉：

> 打开的不是一个工具，而是一本属于自己的咖啡生活杂志。
