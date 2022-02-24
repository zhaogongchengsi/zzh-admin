import type { NextPage } from 'next'
import Head from 'next/head'
import homeStyle from "../styles/home.module.css"
import Menu from "../components/Menus"
import { MoonIcon } from '@heroicons/react/solid'

const Home: NextPage = () => {
  return (
    <div className="w-full h-screen flex flex-col items-center justify-start bg-white dark:bg-gray-800">
      <Head>
        <title>ZZH Blog</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div className="container md:pt-3">
        <Menu />
      </div>
      <div className="w-12 h-12 rounded-full dark:bg-light-600 bg-gray-800 flex items-center justify-center fixed right-4 bottom-4">
        <MoonIcon className="w-6 bg-light-600 dark:bg-gray-800" />
      </div>
    </div>
  )
}

export default Home
