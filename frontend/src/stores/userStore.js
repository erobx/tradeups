import { create } from "zustand"

/*
 * User {
 *  id: number,
 *  username: string,
 *  email: string,
 *  jwt: string,
 *  refresh_token_version: number,
 *  avatarSrc: string,
 *  balance: number,
 * }
 */

const useUserStore = create((set) => ({
    user: {},
    setUser: (user) => set(() => ({ user })),
    setBalance: (balance) => set((state) => ({  
        user: {
            ...state.user,
            balance
        }
    })),
}))

const useUser = () => {
    const { user, setUser, setBalance } = useUserStore()
    return { user, setUser, setBalance }
}

export default useUser
