package Generator

type Production struct {
	left string
	right string
}

type Grammar map[string][]string
var MainGrammar = Grammar{
	"BODY":       	{ "BODY EXPRESSION", "EXPRESSION BODY", "EXPRESSION", "BODY CONDITION", "BODY LOOP_"},
	"EXPRESSION": 	{"EXPRESSION EXPRESSION", "turnRight(ANGLE);", "turnLeft(ANGLE);",
					"turnGunRight(ANGLE);", "turnGunRight(ANGLE);", "fire(POWER);", "doNothing();", "back(PIXELS);", "ahead(PIXELS);"},
	"CONDITION":  	{"if(COND_EXPR){INSIDE_B}"},
	"LOOP_":       	{"while(COND_EXPR){INSIDE_B}"},
	"INSIDE_B":		{"EXPRESSION INSIDE_B", "INSIDE_B EXPRESSION", "EXPRESSION"},
	"COND_EXPR" : 	{"getGunHeat() COND_OPERATOR_EQ 0", "getEnergy() COND_OPERATOR_NEQ ENERGY_LVL"  },
	"ENERGY_LVL" : 	{"20", "50", "random(1,100)"},
	"COND_OPERATOR_EQ" : { "==", "!="},
	"COND_OPERATOR_NEQ" : {">", "<", ">=", "<="},

	//on scanned robot
	"SCANNED" : 		{"SCAN_COND", "SCAN_WHILE", "FIRE_EXPR", "SCANNED EXPRESSION", "EXPRESSION SCANNED"},
	"FIRE_EXPR" :		{"fire(POWER);", "fire(POWER);turnRight(ANGLE);", "fire(POWER);turnLeft(ANGLE);", "fire(POWER); FIRE_EXPR", "turnGunRight(SLIGHT_ANGL); fire(POWER);", "turnGunRight(SLIGHT_ANGL); fire(POWER);"},
	"INSIDE_S" : 		{"FIRE_EXPR", "SCANNED FIRE_EXPR", "FIRE_EXPR SCANNED", "EXPRESSION FIRE_EXPR"},
	"SCAN_WHILE" : 		{"while(SCAN_LOOPCOND){INSIDE_S}"},
	"SCAN_LOOPCOND" : 	{"e.getBearing() COND_OP_DIFF ANGLE", "e.getDistance() COND_OP_DIFF PIXELS", "e.getEnergy() COND_OP_DIFF ENERGY_LVL",
						"e.getEnergy() COND_OP_DIFF getEnergy()" },
	"COND_OP_DIFF" : 	{">", "<", "!="},
	"COND_OP_NDIFF" : 	{">=", "<=", "=="},
	"SCAN_COND" : 		{"if(SCAN_C_EXPR){INSIDE_S}"},
	"SCAN_C_EXPR" : 	{"e.getBearing() COND_OP_NDIFF ANGLE", "e.getDistance() COND_OP_NDIFF PIXELS", "e.getEnergy() COND_OPERATOR_NEQ ENERGY_LVL",
						"e.getEnergy() COND_OPERATOR_NEQ getEnergy()" },

	//for hit wall event
	"ON_HIT_WALL" : 	{ "", "back(PIXELS);", "ahead(PIXELS);", "turnRight(ANGLE);", "turnLeft(ANGLE);", "turnRight(180); ahead(PIXELS);", "turnLeft(180); ahead(PIXELS);" },
	"PIXELS" : 			{ "25", "50", "75", "100", "150", "200", "250", "random(5,400)" },
	"ANGLE" : 			{ "30", "60", "90", "180", "-30", "-60", "-90", "random(-180,180)", "SLIGHT_ANGL" },
	"POWER" : 			{ "1", "2", "3", "4", "5", "random(1,5)" },

	//for hit by bullet event
	"ON_BULLET_HIT" : { "", "EXPRESSION", "ahead(PIXELS);", "back(PIXELS);", "turnGunRight(e.getBearing()); fire(POWER);" , "turnGunRight(e.getBearing() + ANGLE); EXPRESSION",
						"turnRight(ANGLE); ahead(PIXELS);", "turnLeft(ANGLE); ahead(PIXELS);", "turnRight(90-e.getBearing());", "turnLeft(90-e.getBearing());",
						"turnRight(90-e.getBearing()); EXPRESSION", "turnLeft(90-e.getBearing()); EXPRESSION" },

	//for hit robot event
	"ON_HIT_ROBOT" : 	{"", "EXPRESSION", "if(e.isMyFault()){HIT_ROBOT_EXPR}", "HIT_ROBOT_EXPR", "RUNAWAY", "if(e.getEnergy() < getEnergy()) {HIT_ROBOT_EXPR} else {RUNAWAY}",
						"if(!e.isMyFault()){RUNAWAY}"},
	"SLIGHT_ANGL" : 	{"random(-20,20)"},
	"HIT_ROBOT_EXPR" : 	{"fire(POWER);", "fire(POWER); HIT_ROBOT_EXPR", "fire(POWER); RUNAWAY", "turnGunRight(e.getBearing()); fire(POWER);",
						"turnGunRight(e.getBearing() + SLIGHT_ANGL); fire(POWER);"},
	"RUNAWAY" : 		{"turnRight(ANGLE); ahead(PIXELS);", "turnLeft(ANGLE); ahead(PIXELS);", "back(PIXELS);", "turnRight(e.getBearing()-90); ahead(PIXELS);",
						"turnRight(e.getBearing()+90); ahead(PIXELS);"  },
}

var simpleGrammar = Grammar{
	"BODY":          {"EXPRESSION BODY", "EXPRESSION", "CONDITION EXPRESSION", "LOOP EXPRESSION"},
	"SCANNED":       {"fire(POWER);"},
	"POWER":         {"1", "2", "5"},
	"CONDITION":     {"if(true){BODY}"},
	"LOOP":          {"while(getEnergy()>50){BODY}"},
	"EXPRESSION":    {"fire(1);", "turnLeft(50);", "turnRight(50);", "ahead(25);", "back(25);", "EXPRESSION EXPRESSION"},
	"ON_HIT_WALL":   {""},
	"ON_BULLET_HIT": {""},
	"ON_HIT_ROBOT":  {"", "fire(POWER);"},
}

/*var StrategyGrammar = Grammar {
	"BODY" : {},
	"SCANNED" : {},
	"ON_HIT_WALL" : {},
	"ON_BULLET_HIT" : {},
	"ON_HIT_ROBOT" : {},

	"GO_IN_CIRCLES" : {"turnRight(30); ahead(50);", "turnLeft(30); ahead(50);"},
}*/
