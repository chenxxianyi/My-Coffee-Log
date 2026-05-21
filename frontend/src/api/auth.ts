import request from './request'

export interface AuthResponse {
  token: string
  user: {
    id: number
    email: string
    nickname: string
    avatar_url: string
  }
}

export function login(email: string, password: string): Promise<AuthResponse> {
  return request.post('/auth/login', { email, password })
}

export function register(email: string, password: string, nickname: string): Promise<AuthResponse> {
  return request.post('/auth/register', { email, password, nickname })
}

export function getCurrentUser() {
  return request.get('/users/me')
}

export function updateUser(data: { nickname?: string; avatar_url?: string }) {
  return request.put('/users/me', data)
}
