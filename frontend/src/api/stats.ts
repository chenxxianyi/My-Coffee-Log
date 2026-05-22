import request from './request'

export interface FlavorTagItem {
  name: string
  label: string
  count: number
}

export interface StatsOverview {
  month_count: number
  total_count: number
  favorite_coffee_type: string
  favorite_flavor_tag: string
  recent_flavor_tags: FlavorTagItem[]
}

export interface FlavorProfile {
  acidity: number
  bitterness: number
  sweetness: number
  body: number
  aroma: number
  aftertaste: number
}

export interface MonthlyCount {
  month: string
  count: number
}

export async function getStatsOverview(): Promise<StatsOverview> {
  return request.get('/stats/overview') as unknown as Promise<StatsOverview>
}

export async function getFlavorProfile(): Promise<FlavorProfile> {
  return request.get('/stats/flavor-profile') as unknown as Promise<FlavorProfile>
}

export async function getMonthlyStats(): Promise<MonthlyCount[]> {
  return request.get('/stats/monthly') as unknown as Promise<MonthlyCount[]>
}

export interface LifestyleQuoteResponse {
  quote: string
}

export async function getLifestyleQuote(): Promise<LifestyleQuoteResponse> {
  return request.post('/ai/lifestyle-quote') as unknown as Promise<LifestyleQuoteResponse>
}
