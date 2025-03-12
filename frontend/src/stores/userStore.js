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

const useUserIdStore = create((set) => ({
    userId: "",
    setUserId: (userId) => set(() => ({ userId })),
    username: "",
    setUsername: (username) => set(() => ({ username })),
}))

const useUserId = () => {
    const { userId, setUserId } = useUserIdStore()
    return { userId, setUserId }
}

export default useUserId
