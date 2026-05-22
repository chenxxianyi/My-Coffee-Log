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
  flavor_tags: { id: number; name: string; label: string; color: string }[]
  created_at: string
  updated_at: string
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
    flavor_tags: dto.flavor_tags?.map(t => t.name) || []
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
    flavor_tag_ids
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
