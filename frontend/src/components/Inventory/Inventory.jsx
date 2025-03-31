import InventoryItem from "./InventoryItem"
import EmptyItem from "../EmptyItem"
import useInventory from "../../stores/inventoryStore"
import { usePresignedUrls } from "../../hooks/usePresignedUrls"
import { useEffect, useMemo, useState } from "react"
import { rarityOrder } from "../../constants/rarity"
import { wearOrder } from "../../constants/wear"
import { deleteSkin, getInventory } from "../../api/inventory"
import useUser from "../../stores/userStore"
import PageSelector from "./PageSelector"

function Modal({ invId, removeItem }) {
  const { user, setUser, setBalance } = useUser()

  const onClick = async () => {
    const jwt = localStorage.getItem("jwt")
    const res = await deleteSkin(jwt, user.id, invId)
    if (res.status !== 204) {
      return
    }
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
  const { user, setUser } = useUser()
  const { inventory, setInventory, addItem, removeItem } = useInventory()
  const [loading, setLoading] = useState(false)
  const [filter, setFilter] = useState("")
  const [lastInventoryHash, setLastInventoryHash] = useState(0)
  const processedInventory = usePresignedUrls(inventory)
  const [currentPage, setCurrentPage] = useState(1)
  const itemsPerPage = 18

  const createInventoryHash = (items) => {
    return items.map(item => item.id).sort().join('|')
  }

  useEffect(() => {
    const loadInventory = async () => {
      setLoading(true)
      const jwt = localStorage.getItem("jwt")
      try {
        const data = await getInventory(jwt, user.id)

        if (data.skins) {
          const freshInventory = data.skins
          const newHash = createInventoryHash(freshInventory)

          if (newHash !== lastInventoryHash) {
            setInventory(freshInventory)
            setLastInventoryHash(newHash)
          }
        }
      } catch (error) {
        console.error("Failed to fetch inventory:", error)
      } finally {
        setLoading(false)
      }
    }
    loadInventory()
  }, [setInventory, lastInventoryHash])

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
    const label = e.target.getAttribute('aria-label')
    setFilter(label || "")
    setCurrentPage(1)
  }

  const refreshInventory = async () => {
    setLoading(true)
    const jwt = localStorage.getItem("jwt")
    try {
      const data = await getInventory(jwt, user.id)
      if (data.skins) {
        const freshInventory = data.skins
        const newHash = createInventoryHash(freshInventory)
        if (newHash !== lastInventoryHash) {
          setInventory(freshInventory)
          setLastInventoryHash(newHash)
        }
      }
    } catch (error) {
      console.error("Failed to refresh inventory:", error)
    } finally {
      setLoading(false)
    }
  }

  // TODO: allow bulk deletes
  const enterDeleteMode = () => {

  }

  if (loading) {
    return (
      <div className="flex justify-center items-center min-h-screen">
        <div className="loading loading-spinner loading-xl"></div>
      </div>
    )
  }
  
  if (processedInventory.length === 0) {
    return <EmptyItem />
  }

  return (
    <div className="flex flex-col gap-4 lg:flex-row">
      <div className="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-2 flex-grow">
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

      <div className="w-full lg:w-fit lg:ml-30 lg:mb-0">
        <div className="card flex flex-col items-center text-center gap-3 bg-base-200 p-4 w-full lg:w-[14vw] lg:mr-2">
          <h1 className="font-bold text-lg">Settings</h1>
          <form className="filter" onClick={handleFilter}>
            <input className="btn btn-soft btn-square" type="reset" value="×"/>
            <input className="btn btn-soft btn-info" type="radio" name="frameworks" aria-label="Rarity"/>
            <input className="btn btn-soft btn-accent" type="radio" name="frameworks" aria-label="Wear"/>
            <input className="btn btn-soft btn-warning" type="radio" name="frameworks" aria-label="Price"/>
          </form>
          <div className="w-full">
            <button className="btn btn-soft btn-error w-full">Enter delete mode</button>
          </div>
          <div className="w-full">
            <button 
              className="btn btn-soft btn-primary w-full" 
              onClick={refreshInventory}>
              Refresh Inventory
            </button>
          </div>
        </div>
      </div>

      <div className="fixed bottom-4 right-8 z-50">
        <div className="join">
          <button className="join-item btn" onClick={() => paginate(currentPage - 1)} disabled={currentPage === 1}>«</button>
          <div className="join-item btn"> 
            <PageSelector totalPages={totalPages} currentPage={currentPage} setCurrentPage={setCurrentPage} />
          </div>
          <button className="join-item btn" onClick={() => paginate(currentPage + 1)} disabled={currentPage === totalPages}>»</button>
        </div>
      </div>
    </div>
  )
}

export default Inventory
