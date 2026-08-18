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
  'Calm': '平静',
  'Energetic': '愉悦',
  'Reflective': '沉浸',
  'Tired': '疲惫'
}

export const MOOD_ICONS: Record<string, string> = {
  'Calm': 'calm',
  'Energetic': 'energetic',
  'Reflective': 'reflective',
  'Tired': 'tired'
}

export const LIFESTYLE_MOOD_TAGS = [
  { val: 'Calm', label: '平静 Calm', icon: 'calm' },
  { val: 'Focused', label: '专注 Focused', icon: 'focused' },
  { val: 'Tired', label: '疲惫 Tired', icon: 'tired' },
  { val: 'Happy', label: '开心 Happy', icon: 'happy' },
  { val: 'Rainy', label: '阴雨 Rainy', icon: 'rainy' },
  { val: 'Slow', label: '慢活 Slow', icon: 'slow' },
  { val: 'Productive', label: '高效 Productive', icon: 'productive' }
] as const

export const LIFESTYLE_SCENE_TAGS = [
  { val: 'Morning', label: '早晨 Morning', icon: 'morning' },
  { val: 'Office', label: '办公 Office', icon: 'office' },
  { val: 'Weekend', label: '周末 Weekend', icon: 'weekend' },
  { val: 'Cafe', label: '咖啡馆 Cafe', icon: 'cafe' },
  { val: 'Travel', label: '旅行 Travel', icon: 'travel' },
  { val: 'Home', label: '居家 Home', icon: 'home' },
  { val: 'Study', label: '学习 Study', icon: 'study' }
] as const

export const LIFESTYLE_PAIRING_TAGS = [
  { val: 'Book', label: '阅读 Book', icon: 'book' },
  { val: 'Music', label: '音乐 Music', icon: 'music' },
  { val: 'Work', label: '工作 Work', icon: 'work' },
  { val: 'Dessert', label: '甜点 Dessert', icon: 'dessert' },
  { val: 'Alone', label: '独处 Alone', icon: 'alone' },
  { val: 'Friends', label: '聚会 Friends', icon: 'friends' }
] as const

export const coffeeTypeShortLabel = (type: string) => COFFEE_TYPE_SHORT_LABELS[type] ?? type
export const coffeeTypeLabel = (type: string) => COFFEE_TYPE_LABELS[type] ?? type
export const moodLabel = (mood: string) => MOOD_LABELS[mood] ?? mood
export const moodIconName = (mood: string) => MOOD_ICONS[mood] ?? 'coffee'
