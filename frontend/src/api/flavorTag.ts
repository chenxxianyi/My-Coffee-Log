import request from './request'

export interface FlavorTag {
  id: number
  name: string
  label: string
  color: string
}

// Fallback list matching backend seed data (used before API loads)
const FALLBACK_FLAVOR_TAGS: FlavorTag[] = [
  { id: 1, name: 'floral', label: '花香', color: '#D4A5A5' },
  { id: 2, name: 'citrus', label: '柑橘', color: '#F5C156' },
  { id: 3, name: 'berry', label: '莓果', color: '#B5443E' },
  { id: 4, name: 'nutty', label: '坚果', color: '#A67B5B' },
  { id: 5, name: 'chocolate', label: '巧克力', color: '#5C3317' },
  { id: 6, name: 'caramel', label: '焦糖', color: '#D4A017' },
  { id: 7, name: 'creamy', label: '奶油', color: '#F5E6CC' },
  { id: 8, name: 'winey', label: '酒香', color: '#722F37' },
  { id: 9, name: 'smoky', label: '烟熏', color: '#4A4A4A' },
  { id: 10, name: 'herbal', label: '草本', color: '#6B8E23' }
]

let cachedTags: FlavorTag[] | null = null

export async function fetchFlavorTags(): Promise<FlavorTag[]> {
  try {
    const tags = await request.get('/flavor-tags') as unknown as FlavorTag[]
    cachedTags = tags
    return tags
  } catch {
    return cachedTags || FALLBACK_FLAVOR_TAGS
  }
}

export function getFlavorTags(): FlavorTag[] {
  return cachedTags || FALLBACK_FLAVOR_TAGS
}

export function getFlavorTagIdByName(name: string): number | undefined {
  const list = cachedTags || FALLBACK_FLAVOR_TAGS
  return list.find(t => t.name === name)?.id
}

export function getFlavorTagNameById(id: number): string | undefined {
  const list = cachedTags || FALLBACK_FLAVOR_TAGS
  return list.find(t => t.id === id)?.name
}
