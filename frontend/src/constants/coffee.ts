export const COFFEE_TYPE_SHORT_LABELS: Record<string, string> = {
  'Pour Over': '手冲',
  'Latte': '拿铁',
  'Americano': '美式',
  'Cold Brew': '冷萃',
  'Espresso': '浓缩',
  'Dirty': '脏咖啡',
  'Cappuccino': '卡布奇诺',
  'Flat White': '馥芮白'
}

export const COFFEE_TYPE_LABELS: Record<string, string> = {
  'Pour Over': 'Pour Over / 手冲',
  'Latte': 'Latte / 拿铁',
  'Americano': 'Americano / 美式',
  'Cold Brew': 'Cold Brew / 冷萃',
  'Espresso': 'Espresso / 浓缩',
  'Dirty': 'Dirty / 脏咖啡',
  'Cappuccino': 'Cappuccino / 卡布奇诺',
  'Flat White': 'Flat White / 馥芮白'
}

export const MOOD_LABELS: Record<string, string> = {
  'Calm': '😌 平静',
  'Energetic': '⚡ 愉悦',
  'Reflective': '💭 沉浸',
  'Tired': '🥱 疲惫'
}

export const LIFESTYLE_MOOD_TAGS = [
  { val: 'Calm', label: '😌 平静 Calm' },
  { val: 'Focused', label: '🎯 专注 Focused' },
  { val: 'Tired', label: '🥱 疲惫 Tired' },
  { val: 'Happy', label: '😊 开心 Happy' },
  { val: 'Rainy', label: '🌧️ 阴雨 Rainy' },
  { val: 'Slow', label: '🐌 慢活 Slow' },
  { val: 'Productive', label: '🚀 高效 Productive' }
] as const

export const LIFESTYLE_SCENE_TAGS = [
  { val: 'Morning', label: '🌅 早晨 Morning' },
  { val: 'Office', label: '🏢 办公 Office' },
  { val: 'Weekend', label: '🌴 周末 Weekend' },
  { val: 'Cafe', label: '☕ 咖啡馆 Cafe' },
  { val: 'Travel', label: '✈️ 旅行 Travel' },
  { val: 'Home', label: '🏠 居家 Home' },
  { val: 'Study', label: '📚 学习 Study' }
] as const

export const LIFESTYLE_PAIRING_TAGS = [
  { val: 'Book', label: '📖 阅读 Book' },
  { val: 'Music', label: '🎵 音乐 Music' },
  { val: 'Work', label: '💻 工作 Work' },
  { val: 'Dessert', label: '🍰 甜点 Dessert' },
  { val: 'Alone', label: '🧘 独处 Alone' },
  { val: 'Friends', label: '👯 聚会 Friends' }
] as const

export const coffeeTypeShortLabel = (type: string) => COFFEE_TYPE_SHORT_LABELS[type] ?? type
export const coffeeTypeLabel = (type: string) => COFFEE_TYPE_LABELS[type] ?? type
export const moodLabel = (mood: string) => MOOD_LABELS[mood] ?? mood
