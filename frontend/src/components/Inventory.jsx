import InventoryItem from "./InventoryItem"
import EmptyItem from "./EmptyItem"
import useInventory from "../stores/inventoryStore"
import { usePresignedUrls } from "../hooks/usePresignedUrls"
import { useMemo, useState } from "react"
import { rarityOrder } from "../constants/rarity"
import { wearOrder } from "../constants/wear"
import { deleteSkin } from "../api/inventory"
import useUserId from "../stores/userStore"

function Modal({ invId, removeItem }) {
  const { userId, setUserId } = useUserId()

  const onClick = async () => {
    const jwt = localStorage.getItem("jwt")
    const res = await deleteSkin(jwt, userId, invId)
    if (res.status !== 204) {
      return
    }

    // update ui
    removeItem(invId)

    console.log("deleted: ", invId)
  }

  return (
    <dialog id={`modal_${invId}`} className="modal">
      <div className="modal-box max-h-3xl">
        <h3 className="font-bold text-lg mb-1">Details</h3>
        <button className="btn btn-error" onClick={onClick}>Delete skin</button>
      </div>
      <form method="dialog" className="modal-backdrop">
        <button>close</button>
      </form>
    </dialog>
  )
}

function Inventory() {
  const { inventory, setInventory, addItem, removeItem } = useInventory()
  const [loading, setLoading] = useState(false)
  const [filter, setFilter] = useState("")
  const processedInventory = usePresignedUrls(inventory)

  const [currentPage, setCurrentPage] = useState(1)
  const itemsPerPage = 21

  const sortedInventory = useMemo(() => {
    const sorted = [...processedInventory]
    sorted.sort((a, b) => {
      switch (filter) {
        case "Rarity":
          return rarityOrder.indexOf(a.rarity) - rarityOrder.indexOf(b.rarity)
        case "Wear":
          return wearOrder.indexOf(a.wear) - wearOrder.indexOf(b.wear)
        case "Price":
          return b.price - a.price
        default:
          const dateA = new Date(
            a.createdAt.replace(/\//g, '-')
          )
          const dateB = new Date(
            b.createdAt.replace(/\//g, '-')
          )
          return dateA - dateB
      }
    })
    return sorted
  }, [processedInventory, filter])

  const startIndex = (currentPage - 1) * itemsPerPage
  const endIndex = startIndex + itemsPerPage
  const currentItems = sortedInventory.slice(startIndex, endIndex)
  const totalPages = Math.ceil(sortedInventory.length / itemsPerPage)

  const paginate = (pageNumber) => setCurrentPage(pageNumber)
  
  const handleFilter = (e) => {
    const label = e.target.ariaLabel
    setFilter(label || "")
    setCurrentPage(1)
  }

  if (loading) {
    return (
      <div className="flex justify-center items-center min-h-screen">
        <div className="loading loading-spinner loading-xl"></div>
      </div>
    )
  }
  
  if (processedInventory.length === 0) {
    return (
      <div>
        <EmptyItem />
      </div>
    )
  }

  return (
    <div className="flex">
      <div className="grid grid-flow-row lg:grid-cols-7 gap-2 md:grid-cols-2">
        {currentItems.map((item) => (
          <div key={item.id} className="item" onClick={() => document.getElementById(`modal_${item.id}`).showModal()}>
            <InventoryItem 
              name={item.name}
              rarity={item.rarity}
              wear={item.wear}
              price={item.price}
              isStatTrak={item.isStatTrak}
              imgSrc={item.imageSrc}
            />
            <Modal
              invId={item.id}
              removeItem={removeItem}
            />
          </div>
        ))}
      </div>
      <div className="mb-2">
        <form className="filter" onClick={handleFilter}>
          <input className="btn btn-square" type="reset" value="×"/>
          <input className="btn" type="radio" name="frameworks" aria-label="Rarity"/>
          <input className="btn" type="radio" name="frameworks" aria-label="Wear"/>
          <input className="btn" type="radio" name="frameworks" aria-label="Price"/>
        </form>
      </div>
      <div className="fixed bottom-4 right-4 z-50">
        <div className="join">
          <button className="join-item btn" onClick={() => paginate(currentPage - 1)} disabled={currentPage === 1}>«</button>
          <button className="join-item btn">Page {currentPage}</button>
          <button className="join-item btn" onClick={() => paginate(currentPage + 1)} disabled={currentPage === totalPages}>»</button>
        </div>
      </div>
    </div>
  )
}

export default Inventory
