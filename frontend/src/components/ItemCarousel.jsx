import { useCallback, useEffect, useRef, useState } from "react"

const imgWidth = 150
const imgHeight = 150

function ItemCarousel({ skins }) {
  const [activeItem, setActiveItem] = useState(0)

  const carouselRef = useRef(null)

  const scrollItem = () => {
    setActiveItem(prevState => {
      if (skins.length-1 > prevState) {
        return prevState + 1
      } else {
        return 0
      }
    })
  }

  const autoplay = useCallback(() => setInterval(scrollItem, 2000),[])

  useEffect(() => {
    const play = autoplay()
    return () => clearInterval(play)
  }, [autoplay])

  useEffect(() => {
    carouselRef.current?.scroll({left: imgWidth * activeItem})
  }, [activeItem])

  return (
    <div ref={carouselRef} className="carousel carousel-center rounded-box w-1/2">
      {
        skins.map((skin, index) => (
          <div key={index} className="carousel-item">
            <img
              width={imgWidth}
              height={imgHeight}
              src={skin.imageSrc}
              alt={skin.imageSrc}
            />
          </div>
        ))
      }
    </div>
  )
}

export default ItemCarousel
