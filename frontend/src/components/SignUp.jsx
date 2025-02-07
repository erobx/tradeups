import { useState } from "react"
import submitSignup from "../api/auth"

function SignUp() {
    const [username, setUsername] = useState('')
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [confirmPassword, setConfirmPassword] = useState('')
    const [loading, setLoading] = useState(false)

    const handleSubmit = async (e) => {
        e.preventDefault()
        setLoading(true)

        if (password !== confirmPassword) {
            return
        }

        const data = await submitSignup(username, email, password)      
        // { jwt: "" }
        // if success we will have returned a jwt, store in localstorage

        setLoading(false)
    }

    return (
        <div className="flex items-center justify-center">
            <fieldset className="fieldset w-xs bg-base-200 border border-base-300 p-4 rounded-box">
                  <legend className="fieldset-legend">Sign Up</legend>

                  <label className="fieldset-label">Username</label>
                  <input type="username" className="input" placeholder="Username" onChange={(e) => setUsername(e.target.value)} />
                  
                  <label className="fieldset-label">Email</label>
                  <input type="email" className="input" placeholder="Email" onChange={(e) => setEmail(e.target.value)}/>
                  
                  <label className="fieldset-label">Password</label>
                  <input type="password" className="input" placeholder="Password" onChange={(e) => setPassword(e.target.value)}/>

                  <label className="fieldset-label">Confirm Password</label>
                  <input type="password" className="input" placeholder="Confirm Password" onChange={(e) => setConfirmPassword(e.target.value)}/>

                  <button onClick={handleSubmit} className="btn btn-neutral mt-4">Sign Up</button>
            </fieldset>
        </div>
    )
}

export default SignUp
