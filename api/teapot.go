package api
import (
	"math/rand"
	"strings"
)

var teapotList = [...]string{
	`
	(  )   (   )  )
	 ) (   )  (  (
	 ( )  (    ) )
	 _____________
	<_____________> ___
	|             |/ _ \
	|               | | |
	|               |_| |
 ___|             |\___/
/    \___________/    \
\_____________________/
`,
	`
			.------.____
		 .-'       \ ___)
	  .-'         \\\
   .-'        ___  \\)
.-'          /  (\  |)
		 __  \  ( | |
		/  \  \__'| |
	   /    \____).-'
	 .'       /   |
	/     .  /    |
  .'     / \/     |
 /      /   \     |
	   /    /    _|_
	   \   /    /\ /\
		\ /    /__v__\
		 '    |       |
			  |     .#|
			  |#.  .##|
			  |#######|
			  |#######|
`,
	`
												/
											   /
							   xxX###xx       /
								::XXX        /
						 xxXX::::::###XXXXXx/#####
					:::XXXXX::::::XXXXXXXXX/    ####
		 xXXX//::::::://///////:::::::::::/#####    #         ##########
	  XXXXXX//:::::://///xXXXXXXXXXXXXXXX/#    #######      ###   ###
	 XXXX        :://///XXXXXXXXX######X/#######      #   ###    #
	 XXXX        ::////XXXXXXXXX#######/ #     #      ####   #  #
	  XXXX/:     ::////XXXXXXXXXX#####/  #     #########      ##
	   ""XX//::::::////XXXXXXXXXXXXXX/###########     #       #
		   "::::::::////XXXXXXXXXXXX/    #     #     #      ##
				 ::::////XXXXXXXXXX/##################   ###
					 ::::://XXXXXX/#    #     #   #######
						 ::::::::/################
								/
							   /
							  /
`,
}

// Teapot returns the god object
func Teapot() string {
	return strings.Replace(teapotList[rand.Int() % len(teapotList)], "\t", "    ", -1)
}
