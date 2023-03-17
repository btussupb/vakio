let popup = document.querySelector(".popup")
let overlay = document.querySelector(".overlay")
let popupOpen = document.querySelector(".popupOpen")
let closePopup = document.querySelector(".popup svg")

if (popup) {
    popupOpen.addEventListener("click", () => {
        popup.classList.toggle("active");
        overlay.classList.toggle("active");
    })

    closePopup.addEventListener("click", () => {
        popup.classList.remove("active");
        overlay.classList.remove("active");
    })
    overlay.addEventListener("click", () => {
        popup.classList.remove("active");
        overlay.classList.remove("active");
    })
}


document.addEventListener('DOMContentLoaded', () => {
    function tabsMain() {
        let descParentTabs = document.querySelector('.tab-head')
        let btn = document.querySelectorAll('.tablink');
        let block = document.querySelectorAll('.tabcontent');
        btn.forEach((key, index) => {
            key.addEventListener('click', function() {
                block.forEach((item, itemindex) => {
                    item.classList.toggle('active', index === itemindex)
                    item.animate([
                        { opacity: '0' },
                        { opacity: '1' }
                    ], { duration: 200, easing: 'ease-in' })
                });
            });
        });

        descParentTabs.addEventListener('click', (e) => {
            const target = e.target
            if (target.classList.contains('tablink')) {
                btn.forEach(item => {
                    item.classList.remove('active')
                })
            }
            target.classList.add('active')
        });
    }
    tabsMain();
})

// tabs

$('.slider-for').slick({
    slidesToShow: 1,
    slidesToScroll: 1,
    fade: true,
    arrows: true,
    prevArrow: '<button type="button" class="slick_arrow slick-prev"><svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M19.8779 10.1857L1.87205 10.1857M1.87205 10.1857L10.875 1.18271M1.87205 10.1857L10.875 19.1886" stroke="#fff" stroke-width="2.12202"/> </svg></button>',
    nextArrow: '<button type="button" class="slick_arrow slick-next"><svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M19.8779 10.1857L1.87205 10.1857M1.87205 10.1857L10.875 1.18271M1.87205 10.1857L10.875 19.1886" stroke="#fff" stroke-width="2.12202"/> </svg></button>',
    appendArrows: '.slider__arrow',
    asNavFor: '.slider-nav'
});
$('.slider-nav').slick({
    slidesToShow: 3,
    slidesToScroll: 1,
    asNavFor: '.slider-for',
    focusOnSelect: true,
    arrows: false,
});