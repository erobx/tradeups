import { create } from "zustand"

const useAuthStore = create((set) => ({
    loggedIn: false,
    setLoggedIn: (loggedIn) => set(() => ({ loggedIn })),
}))

const useAuth = () => {
    const { loggedIn, setLoggedIn } = useAuthStore()
    return { loggedIn, setLoggedIn }
}

export default useAuth
