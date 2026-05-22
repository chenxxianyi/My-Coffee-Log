import request from './request'
import { CoffeeLogDTO, PaginatedResponse, toCoffeeLog } from './coffeeLog'

export interface CoffeeShop {
  id: number
  user_id: number
  name: string
  address: string
  rating: number
  image_url: string
  visit_count: number
  last_visit_at: string | null
  created_at: string
  updated_at: string
}

export interface CreateCoffeeShopParams {
  name: string
  address?: string
  rating?: number
  image_url?: string
}

export interface UpdateCoffeeShopParams {
  name?: string
  address?: string
  rating?: number
  image_url?: string
}

export async function getCoffeeShops(params?: {
  page?: number
  page_size?: number
  search?: string
}): Promise<PaginatedResponse<CoffeeShop>> {
  return request.get('/coffee-shops', { params }) as unknown as Promise<PaginatedResponse<CoffeeShop>>
}

export async function getCoffeeShopById(id: number): Promise<CoffeeShop> {
  return request.get(`/coffee-shops/${id}`) as unknown as Promise<CoffeeShop>
}

export async function createCoffeeShop(params: CreateCoffeeShopParams): Promise<CoffeeShop> {
  return request.post('/coffee-shops', params) as unknown as Promise<CoffeeShop>
}

export async function updateCoffeeShop(id: number, params: UpdateCoffeeShopParams): Promise<CoffeeShop> {
  return request.put(`/coffee-shops/${id}`, params) as unknown as Promise<CoffeeShop>
}

export async function deleteCoffeeShop(id: number) {
  return request.delete(`/coffee-shops/${id}`)
}

export async function getShopNames(): Promise<string[]> {
  const res = await request.get('/coffee-shops/names') as any
  return res.names || []
}

export async function getShopLogs(shopId: number, params?: {
  page?: number
  page_size?: number
}): Promise<PaginatedResponse<ReturnType<typeof toCoffeeLog>>> {
  const res = await request.get(`/coffee-shops/${shopId}/logs`, { params }) as PaginatedResponse<CoffeeLogDTO>
  return {
    list: res.list.map(toCoffeeLog),
    pagination: res.pagination
  }
}
