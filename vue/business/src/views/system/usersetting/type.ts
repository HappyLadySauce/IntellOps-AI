export interface ModeItem {
  title: string
  icon: string
  subtitle: string
  value?: string
  type: 'mobile' | 'email' | 'gitee' | 'github'|'password'
  jumpMode?: 'link' | 'modal'
  status: boolean
  statusString: string
}
