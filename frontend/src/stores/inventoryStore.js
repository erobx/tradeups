import { create } from "zustand"

const useInventoryStore = create((set) => ({
    inventory: [],
    setInventory: (inventory) => set(() => ({ inventory })),
    addItem: (item) => set((state) => ({ inventory: [...state.inventory, item] })),
    removeItem: (itemId) => set((state) => ({ inventory: state.inventory.filter((item) => item.id !== itemId) })),
}))

const useInventory = () => {
    const { inventory, setInventory, addItem, removeItem } = useInventoryStore()
    return { inventory, setInventory, addItem, removeItem }
}

export default useInventory
