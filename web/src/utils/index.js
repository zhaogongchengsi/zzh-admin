


export function JudgeRequestStatus (state) {
    if(state.success === true && state.code === 200) {
        return function (ok, err) {ok()}
    } else {
        return function (ok, err) {err()}
    }
}