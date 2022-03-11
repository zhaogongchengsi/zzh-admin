import type { NextPage } from 'next'
import React, { useMemo, useState } from 'react';
import Head from 'next/head'
import homeStyle from '../../styles/home.module.css'
import { Nav, Input,Switch, Space, Typography } from '@douyinfe/semi-ui';
import { IconArticle, IconBookmark, IconLayers, IconMoon, IconSearch, IconSetting, IconListView, IconGridView} from '@douyinfe/semi-icons';


const HomeContainer: NextPage = ({children}) => {
    const [layout, setlayout] = useState('list')
    const [DarkTheme, setDarkTheme] = useState(false)
    const { Text } = Typography;


    function isDark () {
      const body = document.body;
      const editor = document.querySelector('.ed-my')
      if (DarkTheme === false) {
        window.document.documentElement.classList.add('dark')
        body.setAttribute('theme-mode', 'dark');
        if (editor != null) {
          editor.classList.add('md-dark')
        } 
      } else {
        window.document.documentElement.classList.remove('dark')
        body.removeAttribute('theme-mode');
        if (editor != null) {
          editor.classList.remove('md-dark')
        }
      }
    }

    const onDarkEvent = (checked:boolean) => {
        setDarkTheme(checked)
        isDark()
    }

    return (
        <>
        <div className={homeStyle.blogHomeContainer}>
        <div className={homeStyle.banner}>
          <div className={homeStyle.homeTitle}>这是标题</div>
          <img src="/bg1.png" alt="banner" />
        </div>
        <div className={homeStyle.navs}>
          <Nav mode="horizontal" className="border-b shadow" >
          <Nav.Header logo={<img src="https://sf6-cdn-tos.douyinstatic.com/obj/eden-cn/ptlz_zlp/ljhwZthlaukjlkulzlp/root-web-sites/webcast_logo.svg" />} text={'ZZH-Blog'} ></Nav.Header>
            <Nav.Item itemKey={'golder'}>
              <Input size='small' prefix={<IconSearch />} />
            </Nav.Item>
            <Nav.Sub itemKey={'ArticleType'} text="文章类型" icon={<IconArticle />}>
                    <Nav.Item itemKey={'golang'} text={'Golang'} />
                    <Nav.Item itemKey={'javascript'} text={'Javascript'} />
                    <Nav.Item itemKey={'Css'} text={'Css'} />
            </Nav.Sub>
            <Nav.Sub itemKey={'Copy'} text="快捷黏贴" icon={<IconBookmark />}>
                    <Nav.Item itemKey={'javascript'} text={'Javascript'} />
                    <Nav.Item itemKey={'Css'} text={'Css'} />
                    <Nav.Item itemKey={'Scss'} text={'Scss'} />
            </Nav.Sub>
            <Nav.Item itemKey={'link'} text={'快捷链接'} icon={<IconLayers />} />
            <Nav.Sub itemKey={'setting'} text="选项设置" icon={<IconSetting />}>
                    <Nav.Item itemKey={'dark'} text={'夜间模式'} >
                      <Space>
                        <IconMoon />
                        <Text>{DarkTheme ? '暗黑模式' : '白色模式'}</Text>
                        <Switch size="small" checked={DarkTheme} onChange={onDarkEvent} />
                      </Space>
                    </Nav.Item>
                    <Nav.Item itemKey={'layui'} text={'布局'} >
                      <Space spacing="loose" align="center">
                        <IconGridView className="hover:text-blue-600" onClick={() =>{setlayout('grid')}}/>
                        <Text>{layout === 'grid' ? '网格布局' : '列表布局'}</Text>
                        <IconListView className="hover:text-blue-600" onClick={() =>{setlayout('list')}}/>
                      </Space>
                    </Nav.Item>
            </Nav.Sub>
          </Nav>
        </div>
        <div className={`layout-${layout}`}>
          {children}
        </div>
      </div>
        </>
    )
}

export default HomeContainer;