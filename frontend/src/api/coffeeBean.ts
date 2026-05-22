import request from './request'
import { PaginatedResponse } from './coffeeLog'

export interface CoffeeBean {
  id: number
  user_id: number
  name: string
  origin: string
  processing_method: string
  roast_level: string
  roaster: string
  image_url: string
  usage_count: number
  created_at: string
  updated_at: string
}

export interface CreateCoffeeBeanParams {
  name: string
  origin?: string
  processing_method?: string
  roast_level?: string
  roaster?: string
  image_url?: string
}

export interface UpdateCoffeeBeanParams {
  name?: string
  origin?: string
  processing_method?: string
  roast_level?: string
  roaster?: string
  image_url?: string
}

export async function getCoffeeBeans(params?: {
  page?: number
  page_size?: number
  search?: string
}): Promise<PaginatedResponse<CoffeeBean>> {
  return request.get('/coffee-beans', { params }) as unknown as Promise<PaginatedResponse<CoffeeBean>>
}

export async function getCoffeeBeanById(id: number): Promise<CoffeeBean> {
  return request.get(`/coffee-beans/${id}`) as unknown as Promise<CoffeeBean>
}

export async function createCoffeeBean(params: CreateCoffeeBeanParams): Promise<CoffeeBean> {
  return request.post('/coffee-beans', params) as unknown as Promise<CoffeeBean>
}

export async function updateCoffeeBean(id: number, params: UpdateCoffeeBeanParams): Promise<CoffeeBean> {
  return request.put(`/coffee-beans/${id}`, params) as unknown as Promise<CoffeeBean>
}

export async function deleteCoffeeBean(id: number) {
  return request.delete(`/coffee-beans/${id}`)
}

export async function getBeanList(): Promise<CoffeeBean[]> {
  const res = await request.get('/coffee-beans/list') as any
  return res.beans || []
}
