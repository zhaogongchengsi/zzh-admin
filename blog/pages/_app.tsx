import '../styles/globals.css'
import type { AppProps } from 'next/app'
import Home from '../components/Home/Home'
function MyApp({ Component, pageProps }: AppProps) {
  return (
    <>
      <div className="w-full">
        <Home>
          <Component {...pageProps} />
        </Home>
      </div>
    </>
  )
}

export default MyApp
