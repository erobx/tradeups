import { create } from "zustand"

/*
 * User {
 *  id: number,
 *  username: string,
 * }
 */

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
