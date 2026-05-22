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

export interface PersonalityTag {
  slug: string
  title: string
  subtitle: string
  description: string
  icon: string
}

export interface PersonalityResponse {
  personalities: PersonalityTag[]
}

export async function getLifestyleQuote(): Promise<LifestyleQuoteResponse> {
  return request.post('/ai/lifestyle-quote') as unknown as Promise<LifestyleQuoteResponse>
}

export async function getPersonality(): Promise<PersonalityResponse> {
  return request.get('/stats/personality') as unknown as Promise<PersonalityResponse>
}

// ---- Monthly Review ----

export interface MonthlyReviewFlavorTag {
  name: string
  label: string
  count: number
}

export interface MonthlyReviewCoffeeType {
  coffee_type: string
  count: number
}

export interface MonthlyReviewCoffeeName {
  coffee_name: string
  count: number
}

export interface MonthlyReviewLifestyleTag {
  tag: string
  count: number
}

export interface MonthlyReviewFlavorProfile {
  acidity: number
  bitterness: number
  sweetness: number
  body: number
  aroma: number
  aftertaste: number
}

export interface MonthlyReviewData {
  month: string
  count: number
  favorite_coffee_type: string
  coffee_types: MonthlyReviewCoffeeType[]
  flavor_tags: MonthlyReviewFlavorTag[]
  coffee_names: MonthlyReviewCoffeeName[]
  top_weekday: number | null
  mood_tags: MonthlyReviewLifestyleTag[]
  scene_tags: MonthlyReviewLifestyleTag[]
  pairing_tags: MonthlyReviewLifestyleTag[]
  flavor_profile: MonthlyReviewFlavorProfile | null
  keywords: string[]
  ai_summary: string
}

export interface MonthlyReviewAIResponse {
  summary: string
}

export async function getMonthlyReview(month?: string): Promise<MonthlyReviewData> {
  const params = month ? { month } : {}
  return request.get('/stats/monthly-review', { params }) as unknown as Promise<MonthlyReviewData>
}

export async function getMonthlyReviewAI(month?: string): Promise<MonthlyReviewAIResponse> {
  const params = month ? { month } : {}
  return request.get('/ai/monthly-review', { params }) as unknown as Promise<MonthlyReviewAIResponse>
}

// ---- AI Status ----

export interface AIStatus {
  enabled: boolean
  model: string
}

export async function getAIStatus(): Promise<AIStatus> {
  return request.get('/ai/status') as unknown as Promise<AIStatus>
}

// ---- Share Copy ----

export interface ShareCopyRequest {
  coffee_name: string
  coffee_type?: string
  shop_name?: string
  mood?: string
  notes?: string
}

export interface ShareCopyResponse {
  copy: string
}

export async function generateShareCopy(req: ShareCopyRequest): Promise<ShareCopyResponse> {
  return request.post('/ai/share-copy', req) as unknown as Promise<ShareCopyResponse>
}

// ---- Coffee Profile ----

export interface CoffeeProfileResponse {
  profile: string
}

export async function generateCoffeeProfile(): Promise<CoffeeProfileResponse> {
  return request.post('/ai/coffee-profile') as unknown as Promise<CoffeeProfileResponse>
}

// ---- Preference Insight ----

export interface PreferenceInsightResponse {
  insight: string
}

export async function generatePreferenceInsight(): Promise<PreferenceInsightResponse> {
  return request.post('/ai/preference-insight') as unknown as Promise<PreferenceInsightResponse>
}
