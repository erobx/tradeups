import { create } from "zustand"

/*
 * User {
 *  id: number,
 *  username: string,
 *  email: string,
 *  jwt: string,
 *  refresh_token_version: number,
 *  avatarSrc: string,
 * }
 */

const useUserStore = create((set) => ({
    user: {},
    setUser: (user) => set(() => ({ user }))
}))

const useUser = () => {
    const { user, setUser } = useUserStore()
    return { user, setUser }
}

export default useUser
