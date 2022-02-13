export function formatTime (time: Date | string): string {
    let _time:Date = typeof time !== 'string' ? time : new Date(time)
    let date:Array<number | string> = [
        _time.getFullYear(),
        _time.getMonth() + 1,
        _time.getDate()
    ];
    return date.join("-")
}