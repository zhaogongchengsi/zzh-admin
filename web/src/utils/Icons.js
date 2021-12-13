import * as solid from '@heroicons/vue/solid'
import * as outline from '@heroicons/vue/outline'

function getIcons(icons) {
    let iconArr = []
    for (let icon in icons ){
        iconArr.push({
           iconName: icon,
           component: icons[icon]
       })
    }
    return iconArr;
}

const solids = getIcons(solid),
     outlines = getIcons(outline)

export default {
    solid,
    outline,
    solids,
    outlines
}