import type { NextPage } from 'next'
import Link from 'next/link'
import { MenuIcon, HomeIcon } from '@heroicons/react/solid'
import { useState } from 'react'
import { useRouter } from 'next/router'

const NavBar: NextPage = () => {
  const [menuH, setMnueH] = useState(false)
  const router = useRouter()
  const menusOpen = () => {
    setMnueH(!menuH)
    console.log(router)
  }

  return (
    <>
      <div className="relative flex flex-col text-gray-800 dark:text-gray-100 md:w-80 md:flex-row">
        <div className="block md:hidden" onClick={menusOpen}>
          <MenuIcon className="w-8" />
        </div>
        <div
          className={`absolute md:h-auto right-0 top-8 flex flex-col overflow-hidden bg-gray-300 md:static md:w-auto md:flex-row md:bg-transparent transition-all duration-75 ${menuH?"h-64":"h-0"}`}
        >
          <Link href="/">
            <a className="flex w-24">
                <HomeIcon className="w-4" />
                <span>首页</span>
            </a>
          </Link>
        </div>
      </div>
    </>
  )
}

export default NavBar
