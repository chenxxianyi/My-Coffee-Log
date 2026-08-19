import request from './request'
import { getFlavorTagIdByName } from './flavorTag'
import { getLocalDateString } from '@/utils/date'

export interface CoffeeLogDTO {
  id: number
  user_id: number
  coffee_name: string
  coffee_type: string
  shop_name: string
  location: string
  image_url: string
  drink_date: string
  mood: string
  notes: string
  acidity: number
  bitterness: number
  sweetness: number
  body: number
  aroma: number
  aftertaste: number
  ai_summary: string
  mood_tags: string
  scene_tags: string
  pairing_tags: string
  flavor_tags: { id: number; name: string; label: string; color: string }[]
  bean_id: number | null
  bean: { id: number; name: string; origin: string; processing_method: string; roast_level: string; roaster: string; image_url: string } | null
  brew_ratio: string
  water_temp: string
  grind_size: string
  created_at: string
  updated_at: string
  // Data quality fields (v2)
  record_mode: string
  coffee_name_source: string
  notes_source: string
  shop_source: string
  sensory_recorded: boolean
  source_log_id: number | null
  is_test_data: boolean
}

export interface PaginatedResponse<T> {
  list: T[]
  pagination: {
    page: number
    page_size: number
    total: number
  }
}

// Convert backend DTO to frontend-friendly format
export function toCoffeeLog(dto: CoffeeLogDTO) {
  return {
    id: dto.id,
    coffee_name: dto.coffee_name,
    coffee_type: dto.coffee_type,
    shop_name: dto.shop_name,
    location: dto.location || '',
    image_url: dto.image_url || '',
    drink_date: dto.drink_date ? dto.drink_date.split('T')[0] : '',
    mood: dto.mood || '',
    notes: dto.notes || '',
    acidity: dto.acidity,
    bitterness: dto.bitterness,
    sweetness: dto.sweetness,
    body: dto.body,
    aroma: dto.aroma,
    aftertaste: dto.aftertaste,
    ai_summary: dto.ai_summary || '',
    mood_tags: dto.mood_tags ? JSON.parse(dto.mood_tags) : [],
    scene_tags: dto.scene_tags ? JSON.parse(dto.scene_tags) : [],
    pairing_tags: dto.pairing_tags ? JSON.parse(dto.pairing_tags) : [],
    flavor_tags: dto.flavor_tags?.map(t => t.name) || [],
    bean_id: dto.bean_id || null,
    bean: dto.bean || null,
    brew_ratio: dto.brew_ratio || '',
    water_temp: dto.water_temp || '',
    grind_size: dto.grind_size || '',
    // Data quality fields
    record_mode: dto.record_mode || 'quick',
    coffee_name_source: dto.coffee_name_source || 'empty',
    notes_source: dto.notes_source || 'empty',
    shop_source: dto.shop_source || 'empty',
    sensory_recorded: dto.sensory_recorded || false,
    source_log_id: dto.source_log_id || null,
    is_test_data: dto.is_test_data || false
  }
}

export interface CreateCoffeeLogParams {
  coffee_name: string
  coffee_type: string
  shop_name: string
  location?: string
  image_url: string
  drink_date?: string
  mood: string
  notes: string
  generate_ai?: boolean
  acidity: number
  bitterness: number
  sweetness: number
  body: number
  aroma: number
  aftertaste: number
  flavor_tags: string[]
  mood_tags?: string[]
  scene_tags?: string[]
  pairing_tags?: string[]
  bean_id?: number | null
  bean_name?: string
  brew_ratio?: string
  water_temp?: string
  grind_size?: string
  // Data quality fields
  record_mode?: string
  sensory_recorded?: boolean
  source_log_id?: number | null
}

function toCreatePayload(params: CreateCoffeeLogParams) {
  const flavor_tag_ids = params.flavor_tags
    .map(name => getFlavorTagIdByName(name))
    .filter((id): id is number => id !== undefined)

  return {
    coffee_name: params.coffee_name,
    coffee_type: params.coffee_type,
    shop_name: params.shop_name,
    location: params.location || '',
    image_url: params.image_url,
    drink_date: params.drink_date || getLocalDateString(),
    mood: params.mood,
    notes: params.notes,
    generate_ai: params.generate_ai ?? false,
    acidity: params.acidity,
    bitterness: params.bitterness,
    sweetness: params.sweetness,
    body: params.body,
    aroma: params.aroma,
    aftertaste: params.aftertaste,
    flavor_tag_ids,
    mood_tags: params.mood_tags || [],
    scene_tags: params.scene_tags || [],
    pairing_tags: params.pairing_tags || [],
    bean_id: params.bean_id || null,
    bean_name: params.bean_name || '',
    brew_ratio: params.brew_ratio || '',
    water_temp: params.water_temp || '',
    grind_size: params.grind_size || '',
    // Data quality fields
    record_mode: params.record_mode || 'quick',
    sensory_recorded: params.sensory_recorded ?? false,
    source_log_id: params.source_log_id || null
  }
}

export async function createCoffeeLog(params: CreateCoffeeLogParams) {
  const payload = toCreatePayload(params)
  const dto = await request.post<CoffeeLogDTO>('/coffee-logs', payload)
  return toCoffeeLog(dto as unknown as CoffeeLogDTO)
}

export async function getCoffeeLogs(params?: {
  page?: number
  page_size?: number
  month?: string
  coffee_type?: string
  tag_id?: number
}): Promise<PaginatedResponse<ReturnType<typeof toCoffeeLog>>> {
  const res = await request.get('/coffee-logs', { params })
  const data = res as unknown as PaginatedResponse<CoffeeLogDTO>
  return {
    list: data.list.map(toCoffeeLog),
    pagination: data.pagination
  }
}

export async function getCoffeeLogById(id: number) {
  const dto = await request.get<CoffeeLogDTO>(`/coffee-logs/${id}`)
  return toCoffeeLog(dto as unknown as CoffeeLogDTO)
}

export async function updateCoffeeLog(id: number, params: Partial<CreateCoffeeLogParams>) {
  const payload: Record<string, unknown> = { ...params }
  if (params.flavor_tags) {
    payload.flavor_tag_ids = params.flavor_tags
      .map(name => getFlavorTagIdByName(name))
      .filter((id): id is number => id !== undefined)
    delete payload.flavor_tags
  }
  const dto = await request.put<CoffeeLogDTO>(`/coffee-logs/${id}`, payload)
  return toCoffeeLog(dto as unknown as CoffeeLogDTO)
}

export async function deleteCoffeeLog(id: number) {
  return request.delete(`/coffee-logs/${id}`)
}
