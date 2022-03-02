import '../styles/globals.css'
import type { AppProps } from 'next/app'
import Home from '../components/Home/Home'
import React from 'react'

const Context = React.createContext({theme: 'dark'});


function MyApp({ Component, pageProps }: AppProps) {
  return (
    <>
      <div className="w-full dark:bg-gray-900 bg-gray-100">
        <Home>
          <Component {...pageProps} />
        </Home>
      </div>
    </>
  )
}

export default MyApp
