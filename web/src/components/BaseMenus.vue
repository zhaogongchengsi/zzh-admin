<script>
import { h } from 'vue'
import { ElMenu, ElSubMenu, ElMenuItem, ElIcon} from 'element-plus'
import Icons from "@/utils/Icons.js"
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
    },
    iconTheme:{
      type: String,
      default: "solid"
    }
  },
  render () {
    const iconTheme = this.iconTheme

    function renderIcon (iconStr) {
     const iocnComponent = Icons[iconTheme]
      return h(ElIcon,null, {
        default: () => {
          return h(iocnComponent[iconStr], {
            class: "h-10 w-10 text-Gray-900"
          })
        }
      })
    }

    function menuItem(menu) {
      return h(
          ElMenuItem,
          {
            index:menu.Path
          },
          {
            default: () => {
              const _icon = renderIcon(menu.Icon) || ""
              return h("div", null, [
                _icon,
                h("span", null, menu.Label)
              ])
            }
          }
      )
    }

    const sunMenus = this.menuList.map(function (menu) {
     
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
              return h("div", null, [
                renderIcon(menu.Icon) || "",
                h("span", null, menu.Label)
              ])
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
        onSelect: this.$emit("select"),
        onOpen:this.$emit("open"),
        onClose:this.$emit("close")
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

