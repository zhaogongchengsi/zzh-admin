import type { NextPage } from 'next'
import Logo from './logo'
import NavBar from './navBar'
const Menu: NextPage = () => {

    return (
        <>
            <div className="w-full flex justify-between items-center">
                <Logo />
                <NavBar />
            </div>
        </>
    )
}

export default Menu