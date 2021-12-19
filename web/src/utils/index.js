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


export function DadLookSon(condition, dad, son) {
  const result = []
  dad.forEach(dad => {
    let isDad = condition(dad)
    if (isDad === false) {
      return false
    }
    son.forEach(sonItem => {
      let isSon = isDad(sonItem)
      if (isSon === false) {
        return false
      }
      let _children = DadLookSon(condition, [isSon], son)
      if (_children.length > 0) {
        isSon.children = _children
      }
      result.push(isSon)
    })
  });
  return result
}