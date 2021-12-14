export function JudgeRequestStatus(state) {
  if (state.success === true && state.code === 200) {
    return function (ok, err) {
      ok();
    };
  } else {
    return function (ok, err) {
      err();
    };
  }
}

export function isThereChinese(str) {
  const reg = /[\u4E00-\u9FA5]|[\uFE30-\uFFA0]/gi;
  if (!reg.exec(str)) {
    return false;
  } else {
    return true;
  }
}
