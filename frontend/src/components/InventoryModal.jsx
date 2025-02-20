import { useState } from "react"
import ModalItem from "./ModalItem"
import useInventory from "../stores/inventoryStore"

function InventoryModal({ rarity }) {
  const { inventory, setInventory, addItem, removeItem } = useInventory()
  const [currentPage, setCurrentPage] = useState(1)
  const itemsPerPage = 18

  const startIndex = (currentPage - 1) * itemsPerPage
  const endIndex = startIndex + itemsPerPage
  const currentItems = inventory.slice(startIndex, endIndex)
  const totalPages = Math.ceil(inventory.length / itemsPerPage)

  const paginate = (pageNumber) => setCurrentPage(pageNumber)

  const addSkin = () => {

  }

  return (
    <div>
      {/* Open the modal using document.getElementById('ID').showModal() method */}
      <button className="btn btn-primary" onClick={()=>document.getElementById('modal_1').showModal()}>Add Skin</button>
      <dialog id="modal_1" className="modal">
        <div className="modal-box max-w-7xl max-h-3xl">
          <h3 className="font-bold text-lg mb-1">Showing all available skins...</h3>
          <div className="grid grid-cols-6 grid-rows-3 gap-2">
            {currentItems.map((s, index) => (
              <ModalItem
                key={index}
                name={s.name}
                wear={s.wear}
                price={s.skinPrice}
                isStatTrak={s.isStatTrak}
                imgSrc={s.imageSrc}
              />
            ))}
          </div>
          <div className="join mt-1">
            <button className="join-item btn" onClick={() => paginate(currentPage - 1)} disabled={currentPage === 1}>«</button>
            <button className="join-item btn">Page {currentPage}</button>
            <button className="join-item btn" onClick={() => paginate(currentPage + 1)} disabled={currentPage === totalPages}>»</button>
          </div>
        </div>
        <form method="dialog" className="modal-backdrop">
          <button>close</button>
        </form>
      </dialog>
    </div>
  )
}

export default InventoryModal
