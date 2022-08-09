package render

import "strings"

const (
	DIGIT = 8
	NUMBER_WIDTH = 18
	COLON_WIDTH = 7
	LINES = 11 // all
)

type char struct {
	formatted string
}

func (c *char) generate() string {
	var result string

	for i := 0; i < LINES; i++ {
		for j := 0; j < DIGIT; j++ {
			result += c.slice(i, j)
		}
		
		result += "\n"
	}

	return result
}

func (c *char) slice(i int, j int) string {
	lines := []string{}
	
	switch c.formatted[j : j+1] {
	case "0":
		lines = strings.Split(n_0, "\n")
	case "1":
		lines = strings.Split(n_1, "\n")
	case "2":
		lines = strings.Split(n_2, "\n")
	case "3":
		lines = strings.Split(n_3, "\n")
	case "4":
		lines = strings.Split(n_4, "\n")
	case "5":
		lines = strings.Split(n_5, "\n")
	case "6":
		lines = strings.Split(n_6, "\n")
	case "7":
		lines = strings.Split(n_7, "\n")
	case "8":
		lines = strings.Split(n_8, "\n")
	case "9":
		lines = strings.Split(n_9, "\n")
	case ":":
		lines = strings.Split(colon, "\n")
	}

	return lines[i]
}

// const sample = `
//      :!77!~.        :?!^                       JYYJJ7~.             .!5Y:                 .~7JYYJ7^        .~7JYYYJ7^.  
//    ~YGBGGGGP?:      YBB?  .::.                 J555PGGPJ^          ~5GGY^                ~PGBP55GBGY:     !PGBG555GBG5~ 
//   JGB5!::~JGBP~    !GG5.  !GGY                     .^JGGG^       ^YGB5~                 :GGG!.  .?GBY    !GGP!.   :!GGG~
//  ?BGY.     ~GGG:  :PGG^   !BG5      :JYJ^       :!!!?5GG5.     .?GBP7.         .?YY~    .PGG?^::~YGGJ    7BG5:     ^PGG!
// :PGP:       7GGJ  !GGJ    !GG5      ~GGG!       JGGBBGG5:     .YBGJ:           :PGGJ     ^5GGGGGBGGJ.    .JGBGY??J5GGG5.
// ~GGY        ^GG5  ~BGY^:..7GG5....   .:.        .::^75GGP~    YBG5!7????7~:     .:.    .75GGP5YY5GGPY~     ^7Y5PPP5GGG! 
// ^GG5        ~GGY  .?PGGGGGGGGGPPG!                    ?GG5   ~GGGGPP555PGBGY:         .YBGY~.   .:7PGG7        .. ~GGP. 
//  YGG7      .5GG~    .^~!!!JGGP!!7:                   :YGBY   ~BGY:..    .7GB5.        :GGG:        7GG5           JGGJ  
//  :PBG?:  .~5BG?           !GGY       :^:       ..:~7YGBGJ.   .5BG?:   ..^?GBY   :^^.  .YBGY~.    :7PGG7          :PGG~  
//   :JGBG5YPGG5!            !BB5      !GBB7   .Y5PPGGGPY7^      :JPBGP55PPGGP?.  :GBBJ   .7PGBP5YY5GBG5~           7BB5   
//     :7Y555J~.             ^J?7      :7J?:   .?J?7!~^.           :~7?J??7!^.    .!J?^     .~7JY5YYJ!^             ~7J~   
// `

const n_0 = `
     .^7?7!^.     
   .?PGGGGBG5!    
  :PBGJ^::!5BGY.  
 .5GG!      JGBY  
 !GGJ       .5GG^ 
 ?GG!        JGG! 
 7GG7        YGG~ 
 :GGP:      ~GG5. 
  !GBP~.  .7GBP^  
   ~5GBPYYPBGY:   
    .~?5PPY?^     
`
const n_1 = `
      :7YYY?.     
    .JGBGGGG~     
    .~J?^PGG!     
         5GG!     
         5GG~     
         5GG~     
        .5GG~     
        .PGG~     
        .PGG^     
        .GGG^     
        .???:     
`
const n_2 = `
   ^J??7~:        
   !PPGGBGY~      
   ...:^7PBGJ.    
         .5GB7    
          JGBJ    
        :JGGP^    
     .~JGBGY:     
   :75GBPJ^       
  ?GGG57:         
 :GGGGP555555555: 
 .!??JJJJJJJJJJJ: 
`
const n_3 = `
    ~5YYJ?!^.     
    ~5Y5PGBG57.   
         .!PGB?   
     .~!!7YGGP!   
     ^GGGBBGG!    
     .::^~JPGGJ.  
           ^GGB~  
          .7GGG^  
     .:^!JPBGP~   
  !55PGGGG5J~.    
  ~YJ?7!^:.       
`
const n_4 = `
   :?!^           
   YBB?  .::.     
  !GG5.  !GGY     
 :PGG^   !BG5     
 !GGJ    !GG5     
 ~BGY^:..7GG5.... 
 .?PGGGGGGGGGPPG! 
   .^~!!!JGGP!!7: 
         !GGY     
         !BB5     
         ^J?7     
`
const n_5 = `
   !!!:.........  
  .PGGPPPPPPPPPP^ 
  .PGGJ7????????: 
  :GGG:           
  .PBGJ~^:..      
   ^YPGBGGPPY7:   
     .^~~7J5GBG!  
           .JGB5  
  ...:^~7?YPGG5^  
  JPGGGGGP5J7^    
  !J?7!~^.        
`
const n_6 = `
        .!5Y:     
       ~5GGY^     
     ^YGB5~       
   .?GBP7.        
  .YBGJ:          
  YBG5!7????7~:   
 ~GGGGPP555PGBGY: 
 ~BGY:..    .7GB5.
 .5BG?:   ..^?GBY 
  :JPBGP55PPGGP?. 
    :~7?J??7!^.   
`
const n_7 = `
 .??????????7!~.  
 ^GGGPPPPPPPGGGP^ 
 ^GGP:......^PGG! 
 :55Y.      7GG5. 
           ~GGG~  
          ^PGG7   
         ^PGGJ    
        ^PGG?     
       ^PGG?      
      :PBG!       
      ^7J~        
`
const n_8 = `
     .~7JYYJ7^    
    ~PGBP55GBGY:  
   :GGG!.  .?GBY  
   .PGG?^::~YGGJ  
    ^5GGGGGBGGJ.  
  .75GGP5YY5GGPY~ 
 .YBGY~.   .:7PGG7
 :GGG:        7GG5
 .YBGY~.    :7PGG7
  .7PGBP5YY5GBG5~ 
    .~7JY5YYJ!^   
`
const n_9 = `
    .~7JYYYJ7^.   
   !PGBG555GBG5~  
  !GGP!.   :!GGG~ 
  7BG5:     ^PGG! 
  .JGBGY??J5GGG5. 
    ^7Y5PPP5GGG!  
        .. ~GGP.  
           JGGJ   
          :PGG~   
          7BB5    
          ~7J~    
`
const colon = `
       
       
       
 :JYJ^ 
 ~GGG! 
  .:.  
       
       
  :^:  
 !GBB7 
 :7J?: 
`

// const h_0 = `
                                                     
                                                     
                                                     
                                                     
                                                     
//                          .J^                         
//                          ~@?                         
//                          ~&7                         
//                          ~&7                         
//                          ~@7                         
//                          ~@7                         
//                          !@7                         
//                          ^G~                         
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
// `
// const h_1 = ``
// const h_2 = ``
// const h_3 = ``
// const h_4 = ``
// const h_5 = ``
// const h_6 = ``
// const h_7 = ``
// const h_8 = ``
// const h_9 = ``
// const h_10 = ``
// const h_11 = ``
// const h_12 = ``
// const h_13 = ``
// const h_14 = ``
// const h_15 = ``
// const h_16 = ``
// const h_17 = ``
// const h_18 = ``
// const h_19 = ``
// const h_20 = ``
// const h_21 = ``
// const h_22 = ``
// const h_23 = ``

// const m_0 = ``
// const m_1 = ``
// const m_2 = ``
// const m_3 = ``
// const m_4 = ``
// const m_5 = ``
// const m_6 = ``
// const m_7 = ``
// const m_8 = ``
// const m_9 = ``
// const m_10 = ``
// const m_11 = ``
// const m_12 = ``
// const m_13 = ``
// const m_14 = ``
// const m_15 = ``
// const m_16 = ``
// const m_17 = ``
// const m_18 = ``
// const m_19 = ``
// const m_20 = ``
// const m_21 = ``
// const m_22 = ``
// const m_23 = ``
// const m_24 = ``
// const m_25 = `
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
//                           J~                         
//                           ^P7                        
//                            .YY.                      
//                              ?P:                     
//                               ~P!                    
//                                :5?                   
//                                  Y5.                 
//                                   7P^                
//                                    ^~                
                                                     
                                                     
                                                     
                                                     
// `
// const m_26 = ``
// const m_27 = ``
// const m_28 = ``
// const m_29 = ``
// const m_30 = ``
// const m_31 = ``
// const m_32 = ``
// const m_33 = ``
// const m_34 = ``
// const m_35 = ``
// const m_36 = ``
// const m_37 = ``
// const m_38 = ``
// const m_39 = ``
// const m_40 = ``
// const m_41 = ``
// const m_42 = ``
// const m_43 = ``
// const m_44 = ``
// const m_45 = ``
// const m_46 = ``
// const m_47 = ``
// const m_48 = ``
// const m_49 = ``
// const m_50 = ``
// const m_51 = ``
// const m_52 = ``
// const m_53 = ``
// const m_54 = ``
// const m_55 = ``
// const m_56 = ``
// const m_57 = ``
// const m_58 = ``
// const m_59 = ``

// const s_0 = ``
// const s_1 = ``
// const s_2 = ``
// const s_3 = ``
// const s_4 = ``
// const s_5 = ``
// const s_6 = ``
// const s_7 = ``
// const s_8 = ``
// const s_9 = ``
// const s_10 = ``
// const s_11 = ``
// const s_12 = ``
// const s_13 = ``
// const s_14 = ``
// const s_15 = ``
// const s_16 = ``
// const s_17 = ``
// const s_18 = ``
// const s_19 = ``
// const s_20 = ``
// const s_21 = ``
// const s_22 = ``
// const s_23 = ``
// const s_24 = ``
// const s_25 = ``
// const s_26 = ``
// const s_27 = ``
// const s_28 = ``
// const s_29 = ``
// const s_30 = ``
// const s_31 = ``
// const s_32 = ``
// const s_33 = ``
// const s_34 = ``
// const s_35 = ``
// const s_36 = ``
// const s_37 = ``
// const s_38 = ``
// const s_39 = ``
// const s_40 = ``
// const s_41 = ``
// const s_42 = ``
// const s_43 = ``
// const s_44 = ``
// const s_45 = ``
// const s_46 = ``
// const s_47 = ``
// const s_48 = ``
// const s_49 = ``
// const s_50 = ``
// const s_51 = ``
// const s_52 = ``
// const s_53 = ``
// const s_54 = ``
// const s_55 = ``
// const s_56 = ``
// const s_57 = ``
// const s_58 = ``
// const s_59 = ``

// const circle = `
//                     ..::^^^^^:::..                   
//                .^~~~~~^^:::::^^^~~~~^:.              
//            .^~!~^..                .:^~!~:           
//          :!!^.                         .:~!^.        
//        :7~.                               .^7~.      
//      .7!.                                    ^7^     
//     ^?:                                        !7    
//    ~?.                                          ~?   
//   ^?                                             ~?  
//  .J.                                              7~ 
//  !!                                               .J 
//  J:                                                ?^
//  J                                                 !!
//  J.                                                !!
//  ?:                                                ?^
//  ~7                                               .J 
//   J:                                              ?^ 
//   .J.                                            !!  
//    :?:                                          !!   
//     .7~                                       :7~    
//       ~7:                                   .!7:     
//        .~!^.                              :~!^       
//          .^!~^.                        :~!~:         
//             .^~!~^:.             ..:^~!~:            
//                 .^^~~~~~~^^^^~~~~~~^:.               
//                       ...:::....                     
// `
