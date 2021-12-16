<script>
import { h } from 'vue'
import { ElMenu, ElSubMenu, ElMenuItem } from 'element-plus'
export default {
  name: "BaseMenus",
  props: {
    menuList: {
      type: Array,
      default: function () {
        return []
      }
    },
    mode: {
      type: String,
      default: "vertical"
    },
    collapse: {
      type: Boolean,
      default: false
    },
    ellipsis: {
      type:Boolean,
      default: true
    },
    backgroundColor: {
      tyep: String,
      default: "#ffffff"
    },
    textColor: {
      type: String,
      default: "#303133"
    },
    activeTextColor: {
      type: String,
      default: "#409EFF"
    },
    defaultActive: {
      type: String,
      default: " "
    },
    defaultOpeneds: {
      type:Array,
      default: function () {
        return []
      }
    },
    menuTrigger: {
      type: String,
      default: "hover"
    },
    router: {
      type: Boolean,
      default: false
    },
    collapseTransition: {
      type: Boolean,
      default: true
    }
  },
  render () {
    function renderContextText (context) {
      return h("span", {/* props */}, {default: () => { return context.Label }})
    }

    function menuItem(menu) {
      return h(
          ElMenuItem,
          {
            index:menu.Path
          },
          {
            default: () => { 
              return renderContextText(menu)
            }
          }
      )
    }

    const sunMenus = this.menuList.map(function (menu, index) {
      if (menu.children && menu.children.length > 0) {
        return h(
          ElSubMenu,
          { 
            index:menu.Path
          }, 
          {
            default: () => {
                return menu.children.map((child, i) => {
                  return menuItem(child)
                }, this)
            },
            title: () => {
              return menu.Label
            }
          }
        )
      } else {
        return menuItem(menu)
      }
    })
 
    return h(
      ElMenu,
      {
        mode: this.mode,
        collapse: this.collapse,
        ellipsis: this.ellipsis,
        backgroundColor: this.backgroundColor,
        textColor: this.textColor,
        activeTextColor: this.activeTextColor,
        defaultActive: this.defaultActive,
        defaultOpeneds: this.defaultOpeneds,
        menuTrigger: this.menuTrigger,
        router: this.router,
        collapseTransition: this.collapseTransition,
      },
      {
        default: () => { 
          return sunMenus
        }
      }
    )
  }
}
</script>

