import { useNavigate, Link } from "react-router";

function Navbar() {
    const navigate = useNavigate()

    return (
        <div className="navbar bg-base-300 shawdow-sm">
            <div className="flex-1">
                <Link to="/" className="btn btn-ghost text-xl">TradeUps</Link>
            </div>
            <div className="flex-none">
                <Link to="/signin" className="btn btn-ghost text-xl">Sign Up</Link>
            </div>
        </div>
    )
}

export default Navbar
