export type Player = {
  id: string
  name: string
  price: number
}

export const MY_TEAM_ID = '22087246-01bc-46ad-a9d9-a99a6d734167'

export const PlayersMap: Record<string, string> = {
  'Cristiano Ronaldo': '/img/players/Cristiano Ronaldo.png',
  'De Bruyne': '/img/players/De Bruyne.png',
  'Harry Kane': '/img/players/Harry Kane.png',
  Lewandowski: '/img/players/Lewandowski.png',
  Maguirre: '/img/players/Maguirre.png',
  Messi: '/img/players/Messi.png',
  Neymar: '/img/players/Neymar.png',
  Richarlison: '/img/players/Richarlison.png',
  'Vinicius Junior': '/img/players/Vinicius Junior.png'
}

type ActionType = 'goal' | 'yellow card' | 'red card' | 'assist'

export type Action = {
  player_name: string
  minutes: number
  action: ActionType
  score: number
}

export type Match = {
  id: string
  match_date: string
  team_a: string
  team_b: string
  result: string
  actions: Action[]
}

export const TeamsImagesMap: Record<string, string> = {
  Alemanha: '/img/flags/Alemanha.png',
  Argentina: '/img/flags/Argentina.png',
  Bélgica: '/img/flags/Belgica.png',
  Brasil: '/img/flags/Brasil.png',
  França: '/img/flags/Franca.png',
  Inglaterra: '/img/flags/Inglaterra.png',
  Polônia: '/img/flags/Polonia.png',
  Portugal: '/img/flags/Portugal.png'
}
