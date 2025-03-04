import { create } from "zustand"

const useTradeupsStore = create((set) => ({
    tradeups: [],
    setTradeups: (tradeups) => set(() => ({ tradeups })),
    addTradeup: (t) => set((state) => ({ tradeups: [...state.tradeups, t] })),
    removeTradeup: (id) => set((state) => ({ tradeups: state.tradeups.filter((t) => t.id !== id) })),
}))

const useTradeups = () => {
    const { tradeups, setTradeups, addTradeup, removeTradeup } = useTradeupsStore()
    return { tradeups, setTradeups, addTradeup, removeTradeup }
}

export default useTradeups
