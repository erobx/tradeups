import { create } from "zustand"

const useUserIdStore = create((set) => ({
    userId: "",
    setUserId: (userId) => set(() => ({ userId })),
}))

const useUserId = () => {
    const { userId, setUserId } = useUserIdStore()
    return { userId, setUserId }
}

export default useUserId
