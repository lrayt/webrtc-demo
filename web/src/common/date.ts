import dayjs from 'dayjs'
import 'dayjs/locale/zh-cn'

dayjs.locale('zh-cn')

export const DateFormat: string = 'YYYY-MM-DD HH:mm:ss'

export const NowDate = () => dayjs().format(DateFormat)