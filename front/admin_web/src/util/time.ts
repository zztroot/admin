import moment from 'moment'

export let dateFormat = (date: any): string => {
    return moment(date).format('YYYY-MM-DD')
}

export let dateTimeFormat = (date: any): string => {
    return moment(date).format('YYYY-MM-DD HH:mm')
}