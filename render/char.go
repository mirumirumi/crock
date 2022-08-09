package render

// numbers: 18 columns
// colon  :  7 columns

var sample = `
     :!77!~.        :?!^                       JYYJJ7~.             .!5Y:                 .~7JYYJ7^        .~7JYYYJ7^.  
   ~YGBGGGGP?:      YBB?  .::.                 J555PGGPJ^          ~5GGY^                ~PGBP55GBGY:     !PGBG555GBG5~ 
  JGB5!::~JGBP~    !GG5.  !GGY                     .^JGGG^       ^YGB5~                 :GGG!.  .?GBY    !GGP!.   :!GGG~
 ?BGY.     ~GGG:  :PGG^   !BG5      :JYJ^       :!!!?5GG5.     .?GBP7.         .?YY~    .PGG?^::~YGGJ    7BG5:     ^PGG!
:PGP:       7GGJ  !GGJ    !GG5      ~GGG!       JGGBBGG5:     .YBGJ:           :PGGJ     ^5GGGGGBGGJ.    .JGBGY??J5GGG5.
~GGY        ^GG5  ~BGY^:..7GG5....   .:.        .::^75GGP~    YBG5!7????7~:     .:.    .75GGP5YY5GGPY~     ^7Y5PPP5GGG! 
^GG5        ~GGY  .?PGGGGGGGGGPPG!                    ?GG5   ~GGGGPP555PGBGY:         .YBGY~.   .:7PGG7        .. ~GGP. 
 YGG7      .5GG~    .^~!!!JGGP!!7:                   :YGBY   ~BGY:..    .7GB5.        :GGG:        7GG5           JGGJ  
 :PBG?:  .~5BG?           !GGY       :^:       ..:~7YGBGJ.   .5BG?:   ..^?GBY   :^^.  .YBGY~.    :7PGG7          :PGG~  
  :JGBG5YPGG5!            !BB5      !GBB7   .Y5PPGGGPY7^      :JPBGP55PPGGP?.  :GBBJ   .7PGBP5YY5GBG5~           7BB5   
    :7Y555J~.             ^J?7      :7J?:   .?J?7!~^.           :~7?J??7!^.    .!J?^     .~7JY5YYJ!^             ~7J~   
`

var n_0 = `
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
var n_1 = `
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
var n_2 = `
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
var n_3 = `
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
var n_4 = `
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
var n_5 = `
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
var n_6 = `
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
var n_7 = `
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
var n_8 = `
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
var n_9 = `
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
var colon = `
       
       
       
 :JYJ^ 
 ~GGG! 
  .:.  
       
       
  :^:  
 !GBB7 
 :7J?: 
`

// var h_0 = `
                                                     
                                                     
                                                     
                                                     
                                                     
//                          .J^                         
//                          ~@?                         
//                          ~&7                         
//                          ~&7                         
//                          ~@7                         
//                          ~@7                         
//                          !@7                         
//                          ^G~                         
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
// `
// var h_1 = ``
// var h_2 = ``
// var h_3 = ``
// var h_4 = ``
// var h_5 = ``
// var h_6 = ``
// var h_7 = ``
// var h_8 = ``
// var h_9 = ``
// var h_10 = ``
// var h_11 = ``
// var h_12 = ``
// var h_13 = ``
// var h_14 = ``
// var h_15 = ``
// var h_16 = ``
// var h_17 = ``
// var h_18 = ``
// var h_19 = ``
// var h_20 = ``
// var h_21 = ``
// var h_22 = ``
// var h_23 = ``

// var m_0 = ``
// var m_1 = ``
// var m_2 = ``
// var m_3 = ``
// var m_4 = ``
// var m_5 = ``
// var m_6 = ``
// var m_7 = ``
// var m_8 = ``
// var m_9 = ``
// var m_10 = ``
// var m_11 = ``
// var m_12 = ``
// var m_13 = ``
// var m_14 = ``
// var m_15 = ``
// var m_16 = ``
// var m_17 = ``
// var m_18 = ``
// var m_19 = ``
// var m_20 = ``
// var m_21 = ``
// var m_22 = ``
// var m_23 = ``
// var m_24 = ``
// var m_25 = `
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
                                                     
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
// var m_26 = ``
// var m_27 = ``
// var m_28 = ``
// var m_29 = ``
// var m_30 = ``
// var m_31 = ``
// var m_32 = ``
// var m_33 = ``
// var m_34 = ``
// var m_35 = ``
// var m_36 = ``
// var m_37 = ``
// var m_38 = ``
// var m_39 = ``
// var m_40 = ``
// var m_41 = ``
// var m_42 = ``
// var m_43 = ``
// var m_44 = ``
// var m_45 = ``
// var m_46 = ``
// var m_47 = ``
// var m_48 = ``
// var m_49 = ``
// var m_50 = ``
// var m_51 = ``
// var m_52 = ``
// var m_53 = ``
// var m_54 = ``
// var m_55 = ``
// var m_56 = ``
// var m_57 = ``
// var m_58 = ``
// var m_59 = ``

// var s_0 = ``
// var s_1 = ``
// var s_2 = ``
// var s_3 = ``
// var s_4 = ``
// var s_5 = ``
// var s_6 = ``
// var s_7 = ``
// var s_8 = ``
// var s_9 = ``
// var s_10 = ``
// var s_11 = ``
// var s_12 = ``
// var s_13 = ``
// var s_14 = ``
// var s_15 = ``
// var s_16 = ``
// var s_17 = ``
// var s_18 = ``
// var s_19 = ``
// var s_20 = ``
// var s_21 = ``
// var s_22 = ``
// var s_23 = ``
// var s_24 = ``
// var s_25 = ``
// var s_26 = ``
// var s_27 = ``
// var s_28 = ``
// var s_29 = ``
// var s_30 = ``
// var s_31 = ``
// var s_32 = ``
// var s_33 = ``
// var s_34 = ``
// var s_35 = ``
// var s_36 = ``
// var s_37 = ``
// var s_38 = ``
// var s_39 = ``
// var s_40 = ``
// var s_41 = ``
// var s_42 = ``
// var s_43 = ``
// var s_44 = ``
// var s_45 = ``
// var s_46 = ``
// var s_47 = ``
// var s_48 = ``
// var s_49 = ``
// var s_50 = ``
// var s_51 = ``
// var s_52 = ``
// var s_53 = ``
// var s_54 = ``
// var s_55 = ``
// var s_56 = ``
// var s_57 = ``
// var s_58 = ``
// var s_59 = ``

// var circle = `
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
