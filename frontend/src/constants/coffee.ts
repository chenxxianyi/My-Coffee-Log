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

export const coffeeTypeShortLabel = (type: string) => COFFEE_TYPE_SHORT_LABELS[type] ?? type
export const coffeeTypeLabel = (type: string) => COFFEE_TYPE_LABELS[type] ?? type
export const moodLabel = (mood: string) => MOOD_LABELS[mood] ?? mood
