import { Link } from "react-router"
import useAuth from "../stores/authStore"

function Navbar() {
    const { loggedIn, setLoggedIn } = useAuth()
    return (
        <div className="navbar bg-base-300 shawdow-sm">
            <div className="flex-1">
                <Link to="/" className="btn btn-ghost text-xl">TradeUps</Link>
            </div>
            {!loggedIn && (<div className="flex-none">
                <Link to="/login" className="btn btn-ghost text-xl">Login</Link>
            </div>)}
        </div>
    )
}

export default Navbar
