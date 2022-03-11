export function formatTime (time: Date | string): string {
    let _time:Date = typeof time !== 'string' ? time : new Date(time)
    let date:Array<number | string> = [
        _time.getFullYear(),
        _time.getMonth() + 1,
        _time.getDate()
    ];
    return date.join("-")
}

export function randomHexColor() {
 var hex = Math.floor(Math.random() * 16777216).toString(16);
 while (hex.length < 6) {
  hex = '0' + hex;
 }
 return '#' + hex;
}