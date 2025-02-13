import { useState } from "react"
import ModalItem from "./ModalItem"

function InventoryModal() {
  // get allowed skins
  const inventorySkins = [
    {id: 0, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: true, imgSrc: "/aug-wings.png"},
    {id: 1, name: "SG 553 | Tornado", wear: "Well-Worn", price: 5.12, isStatTrak: true, imgSrc: "/sg-tornado.png"},
    {id: 2, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: true, imgSrc: "/aug-wings.png"},
    {id: 3, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: false, imgSrc: "/aug-wings.png"},
    {id: 4, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: false, imgSrc: "/aug-wings.png"},
    {id: 5, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: false, imgSrc: "/aug-wings.png"},
    {id: 6, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: false, imgSrc: "/aug-wings.png"},
    {id: 7, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: false, imgSrc: "/aug-wings.png"},
    {id: 8, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: false, imgSrc: "/aug-wings.png"},
    {id: 9, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: false, imgSrc: "/aug-wings.png"},
    {id: 10, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: false, imgSrc: "/aug-wings.png"},
    {id: 11, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: false, imgSrc: "/aug-wings.png"},
    {id: 12, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: false, imgSrc: "/aug-wings.png"},
    {id: 13, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: false, imgSrc: "/aug-wings.png"},
    {id: 14, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: false, imgSrc: "/aug-wings.png"},
    {id: 15, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: false, imgSrc: "/aug-wings.png"},
    {id: 16, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: false, imgSrc: "/aug-wings.png"},
    {id: 17, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: false, imgSrc: "/aug-wings.png"},
    {id: 18, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: false, imgSrc: "/aug-wings.png"},
    {id: 19, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: false, imgSrc: "/aug-wings.png"},
    {id: 20, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: false, imgSrc: "/aug-wings.png"},
    {id: 21, name: "AUG | Wings", wear: "Battle-Scarred", price: 10.12, isStatTrak: false, imgSrc: "/aug-wings.png"},
  ]

  const [currentPage, setCurrentPage] = useState(1)
  const itemsPerPage = 18

  const startIndex = (currentPage - 1) * itemsPerPage
  const endIndex = startIndex + itemsPerPage
  const currentItems = inventorySkins.slice(startIndex, endIndex)
  const totalPages = Math.ceil(inventorySkins.length / itemsPerPage)

  const paginate = (pageNumber) => setCurrentPage(pageNumber)

  const addSkin = () => {

  }

  return (
    <div>
      {/* Open the modal using document.getElementById('ID').showModal() method */}
      <button className="btn btn-primary" onClick={()=>document.getElementById('my_modal_2').showModal()}>Add Skin</button>
      <dialog id="my_modal_2" className="modal">
        <div className="modal-box max-w-7xl max-h-3xl">
          <h3 className="font-bold text-lg">Showing all available skins...</h3>
          <div className="grid grid-cols-6 grid-rows-3 gap-2">
            {currentItems.map((s, index) => (
              <ModalItem
                key={index}
                name={s.name}
                wear={s.wear}
                price={s.price}
                isStatTrak={s.isStatTrak}
                imgSrc={s.imgSrc}
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
