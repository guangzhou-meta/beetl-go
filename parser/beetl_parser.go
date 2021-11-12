package parser // BeetlParser

import (
	"fmt"
	"reflect"
	"strconv"
)

import (
	"github.com/guangzhou-meta/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 79, 618,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 3, 2, 7, 2, 104, 10, 2, 12, 2, 14, 2, 107, 11,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 7, 3, 113, 10, 3, 12, 3, 14, 3, 116, 11, 3,
	3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	5, 4, 130, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 139,
	10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 156, 10, 4, 3, 4, 3, 4, 5, 4, 160, 10, 4,
	3, 4, 3, 4, 5, 4, 164, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 5, 4, 187, 10, 4, 3, 4, 3, 4, 3, 4, 5, 4, 192, 10, 4, 3, 5,
	3, 5, 3, 5, 3, 5, 7, 5, 198, 10, 5, 12, 5, 14, 5, 201, 11, 5, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 7, 7, 211, 10, 7, 12, 7, 14, 7,
	214, 11, 7, 3, 7, 5, 7, 217, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 223,
	10, 8, 12, 8, 14, 8, 226, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 5, 10, 235, 10, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 7, 11, 242,
	10, 11, 12, 11, 14, 11, 245, 11, 11, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12,
	251, 10, 12, 3, 12, 3, 12, 7, 12, 255, 10, 12, 12, 12, 14, 12, 258, 11,
	12, 3, 12, 5, 12, 261, 10, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13,
	7, 13, 269, 10, 13, 12, 13, 14, 13, 272, 11, 13, 3, 13, 3, 13, 7, 13, 276,
	10, 13, 12, 13, 14, 13, 279, 11, 13, 3, 14, 3, 14, 3, 14, 7, 14, 284, 10,
	14, 12, 14, 14, 14, 287, 11, 14, 3, 15, 3, 15, 3, 15, 7, 15, 292, 10, 15,
	12, 15, 14, 15, 295, 11, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16,
	302, 10, 16, 3, 17, 3, 17, 7, 17, 306, 10, 17, 12, 17, 14, 17, 309, 11,
	17, 3, 17, 3, 17, 3, 18, 6, 18, 314, 10, 18, 13, 18, 14, 18, 315, 3, 18,
	7, 18, 319, 10, 18, 12, 18, 14, 18, 322, 11, 18, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 5, 19, 330, 10, 19, 3, 20, 3, 20, 5, 20, 334, 10, 20,
	3, 21, 5, 21, 337, 10, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 5, 22, 344,
	10, 22, 3, 22, 3, 22, 5, 22, 348, 10, 22, 3, 22, 3, 22, 5, 22, 352, 10,
	22, 3, 23, 3, 23, 3, 23, 5, 23, 357, 10, 23, 3, 24, 3, 24, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 7, 26, 368, 10, 26, 12, 26, 14, 26,
	371, 11, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3,
	28, 3, 28, 5, 28, 383, 10, 28, 3, 29, 3, 29, 3, 29, 5, 29, 388, 10, 29,
	3, 30, 3, 30, 3, 30, 5, 30, 393, 10, 30, 3, 30, 5, 30, 396, 10, 30, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 5, 33, 425, 10, 33, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 5, 33, 445, 10, 33, 3, 33, 5, 33,
	448, 10, 33, 3, 33, 5, 33, 451, 10, 33, 7, 33, 453, 10, 33, 12, 33, 14,
	33, 456, 11, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 7, 35, 464,
	10, 35, 12, 35, 14, 35, 467, 11, 35, 3, 35, 5, 35, 470, 10, 35, 3, 36,
	3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	37, 3, 37, 5, 37, 485, 10, 37, 3, 38, 3, 38, 5, 38, 489, 10, 38, 3, 39,
	3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 5, 39, 501,
	10, 39, 3, 40, 3, 40, 3, 40, 5, 40, 506, 10, 40, 3, 40, 3, 40, 7, 40, 510,
	10, 40, 12, 40, 14, 40, 513, 11, 40, 3, 40, 5, 40, 516, 10, 40, 3, 41,
	3, 41, 3, 41, 5, 41, 521, 10, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3,
	42, 7, 42, 529, 10, 42, 12, 42, 14, 42, 532, 11, 42, 3, 43, 3, 43, 3, 43,
	3, 43, 3, 43, 7, 43, 539, 10, 43, 12, 43, 14, 43, 542, 11, 43, 3, 44, 3,
	44, 3, 44, 3, 44, 7, 44, 548, 10, 44, 12, 44, 14, 44, 551, 11, 44, 5, 44,
	553, 10, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3,
	46, 7, 46, 564, 10, 46, 12, 46, 14, 46, 567, 11, 46, 3, 47, 3, 47, 3, 47,
	3, 47, 7, 47, 573, 10, 47, 12, 47, 14, 47, 576, 11, 47, 5, 47, 578, 10,
	47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 7, 47, 585, 10, 47, 12, 47, 14,
	47, 588, 11, 47, 5, 47, 590, 10, 47, 3, 47, 5, 47, 593, 10, 47, 3, 48,
	3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 5, 48, 601, 10, 48, 3, 49, 3, 49, 3,
	49, 3, 49, 3, 49, 5, 49, 608, 10, 49, 3, 50, 3, 50, 3, 51, 3, 51, 5, 51,
	614, 10, 51, 3, 51, 3, 51, 3, 51, 2, 3, 64, 52, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
	52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86,
	88, 90, 92, 94, 96, 98, 100, 2, 9, 3, 2, 19, 20, 3, 2, 27, 28, 3, 2, 30,
	31, 3, 2, 32, 34, 4, 2, 35, 36, 38, 41, 3, 2, 30, 33, 3, 2, 48, 49, 2,
	677, 2, 105, 3, 2, 2, 2, 4, 110, 3, 2, 2, 2, 6, 191, 3, 2, 2, 2, 8, 193,
	3, 2, 2, 2, 10, 204, 3, 2, 2, 2, 12, 207, 3, 2, 2, 2, 14, 218, 3, 2, 2,
	2, 16, 229, 3, 2, 2, 2, 18, 231, 3, 2, 2, 2, 20, 238, 3, 2, 2, 2, 22, 250,
	3, 2, 2, 2, 24, 264, 3, 2, 2, 2, 26, 280, 3, 2, 2, 2, 28, 288, 3, 2, 2,
	2, 30, 301, 3, 2, 2, 2, 32, 303, 3, 2, 2, 2, 34, 313, 3, 2, 2, 2, 36, 329,
	3, 2, 2, 2, 38, 333, 3, 2, 2, 2, 40, 336, 3, 2, 2, 2, 42, 343, 3, 2, 2,
	2, 44, 356, 3, 2, 2, 2, 46, 358, 3, 2, 2, 2, 48, 360, 3, 2, 2, 2, 50, 364,
	3, 2, 2, 2, 52, 372, 3, 2, 2, 2, 54, 382, 3, 2, 2, 2, 56, 384, 3, 2, 2,
	2, 58, 395, 3, 2, 2, 2, 60, 397, 3, 2, 2, 2, 62, 401, 3, 2, 2, 2, 64, 424,
	3, 2, 2, 2, 66, 457, 3, 2, 2, 2, 68, 461, 3, 2, 2, 2, 70, 471, 3, 2, 2,
	2, 72, 484, 3, 2, 2, 2, 74, 486, 3, 2, 2, 2, 76, 500, 3, 2, 2, 2, 78, 502,
	3, 2, 2, 2, 80, 517, 3, 2, 2, 2, 82, 525, 3, 2, 2, 2, 84, 533, 3, 2, 2,
	2, 86, 543, 3, 2, 2, 2, 88, 556, 3, 2, 2, 2, 90, 560, 3, 2, 2, 2, 92, 592,
	3, 2, 2, 2, 94, 600, 3, 2, 2, 2, 96, 607, 3, 2, 2, 2, 98, 609, 3, 2, 2,
	2, 100, 611, 3, 2, 2, 2, 102, 104, 5, 6, 4, 2, 103, 102, 3, 2, 2, 2, 104,
	107, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 108,
	3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 108, 109, 7, 2, 2, 3, 109, 3, 3, 2, 2,
	2, 110, 114, 7, 21, 2, 2, 111, 113, 5, 6, 4, 2, 112, 111, 3, 2, 2, 2, 113,
	116, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 117,
	3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 117, 118, 7, 22, 2, 2, 118, 5, 3, 2,
	2, 2, 119, 192, 5, 4, 3, 2, 120, 192, 5, 54, 28, 2, 121, 192, 5, 60, 31,
	2, 122, 123, 7, 68, 2, 2, 123, 192, 5, 8, 5, 2, 124, 125, 7, 3, 2, 2, 125,
	126, 5, 48, 25, 2, 126, 129, 5, 6, 4, 2, 127, 128, 7, 6, 2, 2, 128, 130,
	5, 6, 4, 2, 129, 127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 192, 3, 2,
	2, 2, 131, 132, 7, 4, 2, 2, 132, 133, 7, 23, 2, 2, 133, 134, 5, 38, 20,
	2, 134, 135, 7, 24, 2, 2, 135, 138, 5, 6, 4, 2, 136, 137, 7, 5, 2, 2, 137,
	139, 5, 6, 4, 2, 138, 136, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 192,
	3, 2, 2, 2, 140, 141, 7, 7, 2, 2, 141, 142, 5, 48, 25, 2, 142, 143, 5,
	6, 4, 2, 143, 192, 3, 2, 2, 2, 144, 145, 7, 8, 2, 2, 145, 146, 5, 48, 25,
	2, 146, 147, 5, 32, 17, 2, 147, 192, 3, 2, 2, 2, 148, 149, 7, 9, 2, 2,
	149, 192, 5, 22, 12, 2, 150, 151, 7, 17, 2, 2, 151, 159, 5, 4, 3, 2, 152,
	153, 7, 18, 2, 2, 153, 155, 7, 23, 2, 2, 154, 156, 7, 64, 2, 2, 155, 154,
	3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 158, 7, 24,
	2, 2, 158, 160, 5, 4, 3, 2, 159, 152, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2,
	160, 192, 3, 2, 2, 2, 161, 163, 7, 10, 2, 2, 162, 164, 5, 64, 33, 2, 163,
	162, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 192,
	7, 50, 2, 2, 166, 167, 7, 11, 2, 2, 167, 192, 7, 50, 2, 2, 168, 169, 7,
	12, 2, 2, 169, 192, 7, 50, 2, 2, 170, 171, 7, 13, 2, 2, 171, 172, 5, 28,
	15, 2, 172, 173, 7, 50, 2, 2, 173, 192, 3, 2, 2, 2, 174, 175, 7, 14, 2,
	2, 175, 192, 5, 18, 10, 2, 176, 177, 5, 30, 16, 2, 177, 178, 7, 50, 2,
	2, 178, 192, 3, 2, 2, 2, 179, 192, 5, 80, 41, 2, 180, 181, 5, 52, 27, 2,
	181, 182, 7, 50, 2, 2, 182, 192, 3, 2, 2, 2, 183, 184, 9, 2, 2, 2, 184,
	186, 7, 64, 2, 2, 185, 187, 7, 64, 2, 2, 186, 185, 3, 2, 2, 2, 186, 187,
	3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 189, 7, 52, 2, 2, 189, 192, 5, 4,
	3, 2, 190, 192, 7, 50, 2, 2, 191, 119, 3, 2, 2, 2, 191, 120, 3, 2, 2, 2,
	191, 121, 3, 2, 2, 2, 191, 122, 3, 2, 2, 2, 191, 124, 3, 2, 2, 2, 191,
	131, 3, 2, 2, 2, 191, 140, 3, 2, 2, 2, 191, 144, 3, 2, 2, 2, 191, 148,
	3, 2, 2, 2, 191, 150, 3, 2, 2, 2, 191, 161, 3, 2, 2, 2, 191, 166, 3, 2,
	2, 2, 191, 168, 3, 2, 2, 2, 191, 170, 3, 2, 2, 2, 191, 174, 3, 2, 2, 2,
	191, 176, 3, 2, 2, 2, 191, 179, 3, 2, 2, 2, 191, 180, 3, 2, 2, 2, 191,
	183, 3, 2, 2, 2, 191, 190, 3, 2, 2, 2, 192, 7, 3, 2, 2, 2, 193, 194, 7,
	73, 2, 2, 194, 199, 5, 10, 6, 2, 195, 196, 7, 75, 2, 2, 196, 198, 5, 10,
	6, 2, 197, 195, 3, 2, 2, 2, 198, 201, 3, 2, 2, 2, 199, 197, 3, 2, 2, 2,
	199, 200, 3, 2, 2, 2, 200, 202, 3, 2, 2, 2, 201, 199, 3, 2, 2, 2, 202,
	203, 7, 74, 2, 2, 203, 9, 3, 2, 2, 2, 204, 205, 5, 12, 7, 2, 205, 206,
	7, 71, 2, 2, 206, 11, 3, 2, 2, 2, 207, 212, 7, 71, 2, 2, 208, 209, 7, 72,
	2, 2, 209, 211, 7, 71, 2, 2, 210, 208, 3, 2, 2, 2, 211, 214, 3, 2, 2, 2,
	212, 210, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 216, 3, 2, 2, 2, 214,
	212, 3, 2, 2, 2, 215, 217, 5, 14, 8, 2, 216, 215, 3, 2, 2, 2, 216, 217,
	3, 2, 2, 2, 217, 13, 3, 2, 2, 2, 218, 219, 7, 76, 2, 2, 219, 224, 5, 16,
	9, 2, 220, 221, 7, 75, 2, 2, 221, 223, 5, 16, 9, 2, 222, 220, 3, 2, 2,
	2, 223, 226, 3, 2, 2, 2, 224, 222, 3, 2, 2, 2, 224, 225, 3, 2, 2, 2, 225,
	227, 3, 2, 2, 2, 226, 224, 3, 2, 2, 2, 227, 228, 7, 77, 2, 2, 228, 15,
	3, 2, 2, 2, 229, 230, 5, 12, 7, 2, 230, 17, 3, 2, 2, 2, 231, 234, 7, 64,
	2, 2, 232, 235, 7, 63, 2, 2, 233, 235, 5, 20, 11, 2, 234, 232, 3, 2, 2,
	2, 234, 233, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 236, 3, 2, 2, 2, 236,
	237, 7, 50, 2, 2, 237, 19, 3, 2, 2, 2, 238, 243, 7, 64, 2, 2, 239, 240,
	7, 51, 2, 2, 240, 242, 7, 64, 2, 2, 241, 239, 3, 2, 2, 2, 242, 245, 3,
	2, 2, 2, 243, 241, 3, 2, 2, 2, 243, 244, 3, 2, 2, 2, 244, 21, 3, 2, 2,
	2, 245, 243, 3, 2, 2, 2, 246, 247, 7, 23, 2, 2, 247, 248, 5, 64, 33, 2,
	248, 249, 7, 24, 2, 2, 249, 251, 3, 2, 2, 2, 250, 246, 3, 2, 2, 2, 250,
	251, 3, 2, 2, 2, 251, 252, 3, 2, 2, 2, 252, 256, 7, 21, 2, 2, 253, 255,
	5, 24, 13, 2, 254, 253, 3, 2, 2, 2, 255, 258, 3, 2, 2, 2, 256, 254, 3,
	2, 2, 2, 256, 257, 3, 2, 2, 2, 257, 260, 3, 2, 2, 2, 258, 256, 3, 2, 2,
	2, 259, 261, 5, 26, 14, 2, 260, 259, 3, 2, 2, 2, 260, 261, 3, 2, 2, 2,
	261, 262, 3, 2, 2, 2, 262, 263, 7, 22, 2, 2, 263, 23, 3, 2, 2, 2, 264,
	265, 7, 15, 2, 2, 265, 270, 5, 64, 33, 2, 266, 267, 7, 51, 2, 2, 267, 269,
	5, 64, 33, 2, 268, 266, 3, 2, 2, 2, 269, 272, 3, 2, 2, 2, 270, 268, 3,
	2, 2, 2, 270, 271, 3, 2, 2, 2, 271, 273, 3, 2, 2, 2, 272, 270, 3, 2, 2,
	2, 273, 277, 7, 52, 2, 2, 274, 276, 5, 6, 4, 2, 275, 274, 3, 2, 2, 2, 276,
	279, 3, 2, 2, 2, 277, 275, 3, 2, 2, 2, 277, 278, 3, 2, 2, 2, 278, 25, 3,
	2, 2, 2, 279, 277, 3, 2, 2, 2, 280, 281, 7, 16, 2, 2, 281, 285, 7, 52,
	2, 2, 282, 284, 5, 6, 4, 2, 283, 282, 3, 2, 2, 2, 284, 287, 3, 2, 2, 2,
	285, 283, 3, 2, 2, 2, 285, 286, 3, 2, 2, 2, 286, 27, 3, 2, 2, 2, 287, 285,
	3, 2, 2, 2, 288, 293, 5, 30, 16, 2, 289, 290, 7, 51, 2, 2, 290, 292, 5,
	30, 16, 2, 291, 289, 3, 2, 2, 2, 292, 295, 3, 2, 2, 2, 293, 291, 3, 2,
	2, 2, 293, 294, 3, 2, 2, 2, 294, 29, 3, 2, 2, 2, 295, 293, 3, 2, 2, 2,
	296, 302, 7, 64, 2, 2, 297, 302, 5, 66, 34, 2, 298, 299, 7, 64, 2, 2, 299,
	300, 7, 37, 2, 2, 300, 302, 5, 4, 3, 2, 301, 296, 3, 2, 2, 2, 301, 297,
	3, 2, 2, 2, 301, 298, 3, 2, 2, 2, 302, 31, 3, 2, 2, 2, 303, 307, 7, 21,
	2, 2, 304, 306, 5, 34, 18, 2, 305, 304, 3, 2, 2, 2, 306, 309, 3, 2, 2,
	2, 307, 305, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308, 310, 3, 2, 2, 2, 309,
	307, 3, 2, 2, 2, 310, 311, 7, 22, 2, 2, 311, 33, 3, 2, 2, 2, 312, 314,
	5, 36, 19, 2, 313, 312, 3, 2, 2, 2, 314, 315, 3, 2, 2, 2, 315, 313, 3,
	2, 2, 2, 315, 316, 3, 2, 2, 2, 316, 320, 3, 2, 2, 2, 317, 319, 5, 6, 4,
	2, 318, 317, 3, 2, 2, 2, 319, 322, 3, 2, 2, 2, 320, 318, 3, 2, 2, 2, 320,
	321, 3, 2, 2, 2, 321, 35, 3, 2, 2, 2, 322, 320, 3, 2, 2, 2, 323, 324, 7,
	15, 2, 2, 324, 325, 5, 64, 33, 2, 325, 326, 7, 52, 2, 2, 326, 330, 3, 2,
	2, 2, 327, 328, 7, 16, 2, 2, 328, 330, 7, 52, 2, 2, 329, 323, 3, 2, 2,
	2, 329, 327, 3, 2, 2, 2, 330, 37, 3, 2, 2, 2, 331, 334, 5, 40, 21, 2, 332,
	334, 5, 42, 22, 2, 333, 331, 3, 2, 2, 2, 333, 332, 3, 2, 2, 2, 334, 39,
	3, 2, 2, 2, 335, 337, 7, 13, 2, 2, 336, 335, 3, 2, 2, 2, 336, 337, 3, 2,
	2, 2, 337, 338, 3, 2, 2, 2, 338, 339, 7, 64, 2, 2, 339, 340, 7, 54, 2,
	2, 340, 341, 5, 64, 33, 2, 341, 41, 3, 2, 2, 2, 342, 344, 5, 44, 23, 2,
	343, 342, 3, 2, 2, 2, 343, 344, 3, 2, 2, 2, 344, 345, 3, 2, 2, 2, 345,
	347, 7, 50, 2, 2, 346, 348, 5, 64, 33, 2, 347, 346, 3, 2, 2, 2, 347, 348,
	3, 2, 2, 2, 348, 349, 3, 2, 2, 2, 349, 351, 7, 50, 2, 2, 350, 352, 5, 46,
	24, 2, 351, 350, 3, 2, 2, 2, 351, 352, 3, 2, 2, 2, 352, 43, 3, 2, 2, 2,
	353, 354, 7, 13, 2, 2, 354, 357, 5, 28, 15, 2, 355, 357, 5, 50, 26, 2,
	356, 353, 3, 2, 2, 2, 356, 355, 3, 2, 2, 2, 357, 45, 3, 2, 2, 2, 358, 359,
	5, 50, 26, 2, 359, 47, 3, 2, 2, 2, 360, 361, 7, 23, 2, 2, 361, 362, 5,
	64, 33, 2, 362, 363, 7, 24, 2, 2, 363, 49, 3, 2, 2, 2, 364, 369, 5, 64,
	33, 2, 365, 366, 7, 51, 2, 2, 366, 368, 5, 64, 33, 2, 367, 365, 3, 2, 2,
	2, 368, 371, 3, 2, 2, 2, 369, 367, 3, 2, 2, 2, 369, 370, 3, 2, 2, 2, 370,
	51, 3, 2, 2, 2, 371, 369, 3, 2, 2, 2, 372, 373, 5, 64, 33, 2, 373, 53,
	3, 2, 2, 2, 374, 375, 7, 55, 2, 2, 375, 376, 5, 56, 29, 2, 376, 377, 7,
	57, 2, 2, 377, 383, 3, 2, 2, 2, 378, 379, 7, 56, 2, 2, 379, 380, 5, 56,
	29, 2, 380, 381, 7, 57, 2, 2, 381, 383, 3, 2, 2, 2, 382, 374, 3, 2, 2,
	2, 382, 378, 3, 2, 2, 2, 383, 55, 3, 2, 2, 2, 384, 387, 5, 64, 33, 2, 385,
	386, 7, 51, 2, 2, 386, 388, 5, 58, 30, 2, 387, 385, 3, 2, 2, 2, 387, 388,
	3, 2, 2, 2, 388, 57, 3, 2, 2, 2, 389, 392, 5, 82, 42, 2, 390, 391, 7, 37,
	2, 2, 391, 393, 7, 63, 2, 2, 392, 390, 3, 2, 2, 2, 392, 393, 3, 2, 2, 2,
	393, 396, 3, 2, 2, 2, 394, 396, 7, 63, 2, 2, 395, 389, 3, 2, 2, 2, 395,
	394, 3, 2, 2, 2, 396, 59, 3, 2, 2, 2, 397, 398, 7, 58, 2, 2, 398, 399,
	7, 60, 2, 2, 399, 400, 7, 57, 2, 2, 400, 61, 3, 2, 2, 2, 401, 402, 5, 64,
	33, 2, 402, 63, 3, 2, 2, 2, 403, 404, 8, 33, 1, 2, 404, 425, 5, 96, 49,
	2, 405, 406, 7, 46, 2, 2, 406, 425, 5, 84, 43, 2, 407, 425, 5, 78, 40,
	2, 408, 425, 5, 68, 35, 2, 409, 425, 5, 92, 47, 2, 410, 411, 7, 64, 2,
	2, 411, 425, 9, 3, 2, 2, 412, 413, 9, 4, 2, 2, 413, 425, 5, 64, 33, 14,
	414, 415, 9, 3, 2, 2, 415, 425, 7, 64, 2, 2, 416, 417, 7, 42, 2, 2, 417,
	425, 5, 64, 33, 12, 418, 419, 7, 23, 2, 2, 419, 420, 5, 64, 33, 2, 420,
	421, 7, 24, 2, 2, 421, 425, 3, 2, 2, 2, 422, 425, 5, 66, 34, 2, 423, 425,
	5, 70, 36, 2, 424, 403, 3, 2, 2, 2, 424, 405, 3, 2, 2, 2, 424, 407, 3,
	2, 2, 2, 424, 408, 3, 2, 2, 2, 424, 409, 3, 2, 2, 2, 424, 410, 3, 2, 2,
	2, 424, 412, 3, 2, 2, 2, 424, 414, 3, 2, 2, 2, 424, 416, 3, 2, 2, 2, 424,
	418, 3, 2, 2, 2, 424, 422, 3, 2, 2, 2, 424, 423, 3, 2, 2, 2, 425, 454,
	3, 2, 2, 2, 426, 427, 12, 11, 2, 2, 427, 428, 9, 5, 2, 2, 428, 453, 5,
	64, 33, 12, 429, 430, 12, 10, 2, 2, 430, 431, 9, 4, 2, 2, 431, 453, 5,
	64, 33, 11, 432, 433, 12, 9, 2, 2, 433, 434, 9, 6, 2, 2, 434, 453, 5, 64,
	33, 10, 435, 436, 12, 8, 2, 2, 436, 437, 7, 43, 2, 2, 437, 453, 5, 64,
	33, 9, 438, 439, 12, 7, 2, 2, 439, 440, 7, 44, 2, 2, 440, 453, 5, 64, 33,
	8, 441, 442, 12, 6, 2, 2, 442, 444, 7, 45, 2, 2, 443, 445, 5, 64, 33, 2,
	444, 443, 3, 2, 2, 2, 444, 445, 3, 2, 2, 2, 445, 447, 3, 2, 2, 2, 446,
	448, 7, 52, 2, 2, 447, 446, 3, 2, 2, 2, 447, 448, 3, 2, 2, 2, 448, 450,
	3, 2, 2, 2, 449, 451, 5, 64, 33, 2, 450, 449, 3, 2, 2, 2, 450, 451, 3,
	2, 2, 2, 451, 453, 3, 2, 2, 2, 452, 426, 3, 2, 2, 2, 452, 429, 3, 2, 2,
	2, 452, 432, 3, 2, 2, 2, 452, 435, 3, 2, 2, 2, 452, 438, 3, 2, 2, 2, 452,
	441, 3, 2, 2, 2, 453, 456, 3, 2, 2, 2, 454, 452, 3, 2, 2, 2, 454, 455,
	3, 2, 2, 2, 455, 65, 3, 2, 2, 2, 456, 454, 3, 2, 2, 2, 457, 458, 5, 68,
	35, 2, 458, 459, 7, 37, 2, 2, 459, 460, 5, 64, 33, 2, 460, 67, 3, 2, 2,
	2, 461, 465, 7, 64, 2, 2, 462, 464, 5, 72, 37, 2, 463, 462, 3, 2, 2, 2,
	464, 467, 3, 2, 2, 2, 465, 463, 3, 2, 2, 2, 465, 466, 3, 2, 2, 2, 466,
	469, 3, 2, 2, 2, 467, 465, 3, 2, 2, 2, 468, 470, 5, 74, 38, 2, 469, 468,
	3, 2, 2, 2, 469, 470, 3, 2, 2, 2, 470, 69, 3, 2, 2, 2, 471, 472, 5, 68,
	35, 2, 472, 473, 9, 7, 2, 2, 473, 474, 7, 37, 2, 2, 474, 475, 5, 64, 33,
	2, 475, 71, 3, 2, 2, 2, 476, 477, 7, 53, 2, 2, 477, 485, 7, 64, 2, 2, 478,
	479, 7, 29, 2, 2, 479, 485, 7, 64, 2, 2, 480, 481, 7, 25, 2, 2, 481, 482,
	5, 64, 33, 2, 482, 483, 7, 26, 2, 2, 483, 485, 3, 2, 2, 2, 484, 476, 3,
	2, 2, 2, 484, 478, 3, 2, 2, 2, 484, 480, 3, 2, 2, 2, 485, 73, 3, 2, 2,
	2, 486, 488, 7, 42, 2, 2, 487, 489, 5, 76, 39, 2, 488, 487, 3, 2, 2, 2,
	488, 489, 3, 2, 2, 2, 489, 75, 3, 2, 2, 2, 490, 501, 5, 96, 49, 2, 491,
	492, 7, 46, 2, 2, 492, 501, 5, 84, 43, 2, 493, 501, 5, 78, 40, 2, 494,
	501, 5, 92, 47, 2, 495, 501, 5, 68, 35, 2, 496, 497, 7, 23, 2, 2, 497,
	498, 5, 64, 33, 2, 498, 499, 7, 24, 2, 2, 499, 501, 3, 2, 2, 2, 500, 490,
	3, 2, 2, 2, 500, 491, 3, 2, 2, 2, 500, 493, 3, 2, 2, 2, 500, 494, 3, 2,
	2, 2, 500, 495, 3, 2, 2, 2, 500, 496, 3, 2, 2, 2, 501, 77, 3, 2, 2, 2,
	502, 503, 5, 82, 42, 2, 503, 505, 7, 23, 2, 2, 504, 506, 5, 50, 26, 2,
	505, 504, 3, 2, 2, 2, 505, 506, 3, 2, 2, 2, 506, 507, 3, 2, 2, 2, 507,
	511, 7, 24, 2, 2, 508, 510, 5, 72, 37, 2, 509, 508, 3, 2, 2, 2, 510, 513,
	3, 2, 2, 2, 511, 509, 3, 2, 2, 2, 511, 512, 3, 2, 2, 2, 512, 515, 3, 2,
	2, 2, 513, 511, 3, 2, 2, 2, 514, 516, 5, 74, 38, 2, 515, 514, 3, 2, 2,
	2, 515, 516, 3, 2, 2, 2, 516, 79, 3, 2, 2, 2, 517, 518, 5, 82, 42, 2, 518,
	520, 7, 23, 2, 2, 519, 521, 5, 50, 26, 2, 520, 519, 3, 2, 2, 2, 520, 521,
	3, 2, 2, 2, 521, 522, 3, 2, 2, 2, 522, 523, 7, 24, 2, 2, 523, 524, 5, 4,
	3, 2, 524, 81, 3, 2, 2, 2, 525, 530, 7, 64, 2, 2, 526, 527, 7, 53, 2, 2,
	527, 529, 7, 64, 2, 2, 528, 526, 3, 2, 2, 2, 529, 532, 3, 2, 2, 2, 530,
	528, 3, 2, 2, 2, 530, 531, 3, 2, 2, 2, 531, 83, 3, 2, 2, 2, 532, 530, 3,
	2, 2, 2, 533, 540, 5, 90, 46, 2, 534, 539, 5, 86, 44, 2, 535, 539, 5, 88,
	45, 2, 536, 537, 7, 53, 2, 2, 537, 539, 5, 90, 46, 2, 538, 534, 3, 2, 2,
	2, 538, 535, 3, 2, 2, 2, 538, 536, 3, 2, 2, 2, 539, 542, 3, 2, 2, 2, 540,
	538, 3, 2, 2, 2, 540, 541, 3, 2, 2, 2, 541, 85, 3, 2, 2, 2, 542, 540, 3,
	2, 2, 2, 543, 552, 7, 23, 2, 2, 544, 549, 5, 64, 33, 2, 545, 546, 7, 51,
	2, 2, 546, 548, 5, 64, 33, 2, 547, 545, 3, 2, 2, 2, 548, 551, 3, 2, 2,
	2, 549, 547, 3, 2, 2, 2, 549, 550, 3, 2, 2, 2, 550, 553, 3, 2, 2, 2, 551,
	549, 3, 2, 2, 2, 552, 544, 3, 2, 2, 2, 552, 553, 3, 2, 2, 2, 553, 554,
	3, 2, 2, 2, 554, 555, 7, 24, 2, 2, 555, 87, 3, 2, 2, 2, 556, 557, 7, 25,
	2, 2, 557, 558, 5, 64, 33, 2, 558, 559, 7, 26, 2, 2, 559, 89, 3, 2, 2,
	2, 560, 565, 7, 64, 2, 2, 561, 562, 7, 53, 2, 2, 562, 564, 7, 64, 2, 2,
	563, 561, 3, 2, 2, 2, 564, 567, 3, 2, 2, 2, 565, 563, 3, 2, 2, 2, 565,
	566, 3, 2, 2, 2, 566, 91, 3, 2, 2, 2, 567, 565, 3, 2, 2, 2, 568, 577, 7,
	25, 2, 2, 569, 574, 5, 64, 33, 2, 570, 571, 7, 51, 2, 2, 571, 573, 5, 64,
	33, 2, 572, 570, 3, 2, 2, 2, 573, 576, 3, 2, 2, 2, 574, 572, 3, 2, 2, 2,
	574, 575, 3, 2, 2, 2, 575, 578, 3, 2, 2, 2, 576, 574, 3, 2, 2, 2, 577,
	569, 3, 2, 2, 2, 577, 578, 3, 2, 2, 2, 578, 579, 3, 2, 2, 2, 579, 593,
	7, 26, 2, 2, 580, 589, 7, 21, 2, 2, 581, 586, 5, 94, 48, 2, 582, 583, 7,
	51, 2, 2, 583, 585, 5, 94, 48, 2, 584, 582, 3, 2, 2, 2, 585, 588, 3, 2,
	2, 2, 586, 584, 3, 2, 2, 2, 586, 587, 3, 2, 2, 2, 587, 590, 3, 2, 2, 2,
	588, 586, 3, 2, 2, 2, 589, 581, 3, 2, 2, 2, 589, 590, 3, 2, 2, 2, 590,
	591, 3, 2, 2, 2, 591, 593, 7, 22, 2, 2, 592, 568, 3, 2, 2, 2, 592, 580,
	3, 2, 2, 2, 593, 93, 3, 2, 2, 2, 594, 595, 7, 63, 2, 2, 595, 596, 7, 52,
	2, 2, 596, 601, 5, 64, 33, 2, 597, 598, 7, 64, 2, 2, 598, 599, 7, 52, 2,
	2, 599, 601, 5, 64, 33, 2, 600, 594, 3, 2, 2, 2, 600, 597, 3, 2, 2, 2,
	601, 95, 3, 2, 2, 2, 602, 608, 7, 60, 2, 2, 603, 608, 7, 62, 2, 2, 604,
	608, 7, 63, 2, 2, 605, 608, 5, 98, 50, 2, 606, 608, 7, 47, 2, 2, 607, 602,
	3, 2, 2, 2, 607, 603, 3, 2, 2, 2, 607, 604, 3, 2, 2, 2, 607, 605, 3, 2,
	2, 2, 607, 606, 3, 2, 2, 2, 608, 97, 3, 2, 2, 2, 609, 610, 9, 8, 2, 2,
	610, 99, 3, 2, 2, 2, 611, 613, 7, 23, 2, 2, 612, 614, 5, 50, 26, 2, 613,
	612, 3, 2, 2, 2, 613, 614, 3, 2, 2, 2, 614, 615, 3, 2, 2, 2, 615, 616,
	7, 24, 2, 2, 616, 101, 3, 2, 2, 2, 69, 105, 114, 129, 138, 155, 159, 163,
	186, 191, 199, 212, 216, 224, 234, 243, 250, 256, 260, 270, 277, 285, 293,
	301, 307, 315, 320, 329, 333, 336, 343, 347, 351, 356, 369, 382, 387, 392,
	395, 424, 444, 447, 450, 452, 454, 465, 469, 484, 488, 500, 505, 511, 515,
	520, 530, 538, 540, 549, 552, 565, 574, 577, 586, 589, 592, 600, 607, 613,
}
var literalNames = []string{
	"", "'if'", "'for'", "'elsefor'", "'else'", "'while'", "'switch'", "'select'",
	"'return'", "'break'", "'continue'", "'var'", "", "'case'", "'default'",
	"'try'", "'catch'", "'#ajax'", "'#fragment'", "'{'", "'}'", "", "", "'['",
	"']'", "'++'", "'--'", "'.~'", "'+'", "'-'", "'*'", "'/'", "'%'", "'=='",
	"'!='", "'='", "'>='", "", "'<='", "", "'!'", "'&&'", "'||'", "'?'", "'@'",
	"'null'", "'true'", "'false'", "';'", "", "':'", "", "'in'", "'<<'", "'<#'",
	"'>>'", "'<$'",
}
var symbolicNames = []string{
	"", "If", "For", "Elsefor", "Else", "While", "Switch", "Select", "Return",
	"Break", "Continue", "Var", "Directive", "Case", "Default", "Try", "Catch",
	"Ajax", "Fragment", "LEFT_BRACE", "RIGHT_BRACE", "LEFT_PAR", "RIGHT_PAR",
	"LEFT_SQBR", "RIGHT_SQBR", "INCREASE", "DECREASE", "VIRTUAL", "ADD", "MIN",
	"MUlTIP", "DIV", "MOD", "EQUAL", "NOT_EQUAL", "ASSIN", "LARGE_EQUAL", "LARGE",
	"LESS_EQUAL", "LESS", "NOT", "AND", "OR", "QUESTOIN", "AT", "NULL", "TRUE",
	"FALSE", "END", "COMMA", "COLON", "PERIOD", "FOR_IN", "LEFT_TOKEN", "LEFT_TOKEN2",
	"RIGHT_TOKEN", "LEFT_TEXT_TOKEN", "HexLiteral", "DecimalLiteral", "OctalLiteral",
	"FloatingPointLiteral", "StringLiteral", "Identifier", "WS", "LINE_COMMENT",
	"COMMENT_OPEN", "COMMENT_TAG", "COMMENT_END", "ALL_COMMENT_CHAR", "Identifier1",
	"PERIOD1", "LEFT_PAR1", "RIGHT_PAR1", "COMMA1", "LEFT_ANGULAR", "RIGHT_ANGULAR",
	"WS1", "TYPE_END",
}

var ruleNames = []string{
	"prog", "block", "statement", "commentTypeTag", "commentTypeItemTag", "classOrInterfaceType",
	"typeArguments", "typeArgument", "directiveExp", "directiveExpIDList",
	"g_switchStatment", "g_caseStatment", "g_defaultStatment", "varDeclareList",
	"assignMent", "switchBlock", "switchBlockStatementGroup", "switchLabel",
	"forControl", "forInControl", "generalForControl", "forInit", "forUpdate",
	"parExpression", "expressionList", "statementExpression", "textStatment",
	"textVar", "textformat", "constantsTextStatment", "constantExpression",
	"expression", "generalAssignExp", "varRef", "generalAssingSelfExp", "varAttribute",
	"safe_output", "safe_allow_exp", "functionCall", "functionTagCall", "functionNs",
	"nativeCall", "nativeMethod", "nativeArray", "nativeVarRefChain", "json",
	"jsonKeyValue", "literal", "booleanLiteral", "arguments",
}

type BeetlParser struct {
	*antlr.BaseParser
}

// NewBeetlParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *BeetlParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewBeetlParser(input antlr.TokenStream) *BeetlParser {
	this := new(BeetlParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "BeetlParser.g4"

	return this
}

// BeetlParser tokens.
const (
	BeetlParserEOF                  = antlr.TokenEOF
	BeetlParserIf                   = 1
	BeetlParserFor                  = 2
	BeetlParserElsefor              = 3
	BeetlParserElse                 = 4
	BeetlParserWhile                = 5
	BeetlParserSwitch               = 6
	BeetlParserSelect               = 7
	BeetlParserReturn               = 8
	BeetlParserBreak                = 9
	BeetlParserContinue             = 10
	BeetlParserVar                  = 11
	BeetlParserDirective            = 12
	BeetlParserCase                 = 13
	BeetlParserDefault              = 14
	BeetlParserTry                  = 15
	BeetlParserCatch                = 16
	BeetlParserAjax                 = 17
	BeetlParserFragment             = 18
	BeetlParserLEFT_BRACE           = 19
	BeetlParserRIGHT_BRACE          = 20
	BeetlParserLEFT_PAR             = 21
	BeetlParserRIGHT_PAR            = 22
	BeetlParserLEFT_SQBR            = 23
	BeetlParserRIGHT_SQBR           = 24
	BeetlParserINCREASE             = 25
	BeetlParserDECREASE             = 26
	BeetlParserVIRTUAL              = 27
	BeetlParserADD                  = 28
	BeetlParserMIN                  = 29
	BeetlParserMUlTIP               = 30
	BeetlParserDIV                  = 31
	BeetlParserMOD                  = 32
	BeetlParserEQUAL                = 33
	BeetlParserNOT_EQUAL            = 34
	BeetlParserASSIN                = 35
	BeetlParserLARGE_EQUAL          = 36
	BeetlParserLARGE                = 37
	BeetlParserLESS_EQUAL           = 38
	BeetlParserLESS                 = 39
	BeetlParserNOT                  = 40
	BeetlParserAND                  = 41
	BeetlParserOR                   = 42
	BeetlParserQUESTOIN             = 43
	BeetlParserAT                   = 44
	BeetlParserNULL                 = 45
	BeetlParserTRUE                 = 46
	BeetlParserFALSE                = 47
	BeetlParserEND                  = 48
	BeetlParserCOMMA                = 49
	BeetlParserCOLON                = 50
	BeetlParserPERIOD               = 51
	BeetlParserFOR_IN               = 52
	BeetlParserLEFT_TOKEN           = 53
	BeetlParserLEFT_TOKEN2          = 54
	BeetlParserRIGHT_TOKEN          = 55
	BeetlParserLEFT_TEXT_TOKEN      = 56
	BeetlParserHexLiteral           = 57
	BeetlParserDecimalLiteral       = 58
	BeetlParserOctalLiteral         = 59
	BeetlParserFloatingPointLiteral = 60
	BeetlParserStringLiteral        = 61
	BeetlParserIdentifier           = 62
	BeetlParserWS                   = 63
	BeetlParserLINE_COMMENT         = 64
	BeetlParserCOMMENT_OPEN         = 65
	BeetlParserCOMMENT_TAG          = 66
	BeetlParserCOMMENT_END          = 67
	BeetlParserALL_COMMENT_CHAR     = 68
	BeetlParserIdentifier1          = 69
	BeetlParserPERIOD1              = 70
	BeetlParserLEFT_PAR1            = 71
	BeetlParserRIGHT_PAR1           = 72
	BeetlParserCOMMA1               = 73
	BeetlParserLEFT_ANGULAR         = 74
	BeetlParserRIGHT_ANGULAR        = 75
	BeetlParserWS1                  = 76
	BeetlParserTYPE_END             = 77
)

// BeetlParser rules.
const (
	BeetlParserRULE_prog                      = 0
	BeetlParserRULE_block                     = 1
	BeetlParserRULE_statement                 = 2
	BeetlParserRULE_commentTypeTag            = 3
	BeetlParserRULE_commentTypeItemTag        = 4
	BeetlParserRULE_classOrInterfaceType      = 5
	BeetlParserRULE_typeArguments             = 6
	BeetlParserRULE_typeArgument              = 7
	BeetlParserRULE_directiveExp              = 8
	BeetlParserRULE_directiveExpIDList        = 9
	BeetlParserRULE_g_switchStatment          = 10
	BeetlParserRULE_g_caseStatment            = 11
	BeetlParserRULE_g_defaultStatment         = 12
	BeetlParserRULE_varDeclareList            = 13
	BeetlParserRULE_assignMent                = 14
	BeetlParserRULE_switchBlock               = 15
	BeetlParserRULE_switchBlockStatementGroup = 16
	BeetlParserRULE_switchLabel               = 17
	BeetlParserRULE_forControl                = 18
	BeetlParserRULE_forInControl              = 19
	BeetlParserRULE_generalForControl         = 20
	BeetlParserRULE_forInit                   = 21
	BeetlParserRULE_forUpdate                 = 22
	BeetlParserRULE_parExpression             = 23
	BeetlParserRULE_expressionList            = 24
	BeetlParserRULE_statementExpression       = 25
	BeetlParserRULE_textStatment              = 26
	BeetlParserRULE_textVar                   = 27
	BeetlParserRULE_textformat                = 28
	BeetlParserRULE_constantsTextStatment     = 29
	BeetlParserRULE_constantExpression        = 30
	BeetlParserRULE_expression                = 31
	BeetlParserRULE_generalAssignExp          = 32
	BeetlParserRULE_varRef                    = 33
	BeetlParserRULE_generalAssingSelfExp      = 34
	BeetlParserRULE_varAttribute              = 35
	BeetlParserRULE_safe_output               = 36
	BeetlParserRULE_safe_allow_exp            = 37
	BeetlParserRULE_functionCall              = 38
	BeetlParserRULE_functionTagCall           = 39
	BeetlParserRULE_functionNs                = 40
	BeetlParserRULE_nativeCall                = 41
	BeetlParserRULE_nativeMethod              = 42
	BeetlParserRULE_nativeArray               = 43
	BeetlParserRULE_nativeVarRefChain         = 44
	BeetlParserRULE_json                      = 45
	BeetlParserRULE_jsonKeyValue              = 46
	BeetlParserRULE_literal                   = 47
	BeetlParserRULE_booleanLiteral            = 48
	BeetlParserRULE_arguments                 = 49
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(BeetlParserEOF, 0)
}

func (s *ProgContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ProgContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *BeetlParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BeetlParserRULE_prog)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeetlParserIf)|(1<<BeetlParserFor)|(1<<BeetlParserWhile)|(1<<BeetlParserSwitch)|(1<<BeetlParserSelect)|(1<<BeetlParserReturn)|(1<<BeetlParserBreak)|(1<<BeetlParserContinue)|(1<<BeetlParserVar)|(1<<BeetlParserDirective)|(1<<BeetlParserTry)|(1<<BeetlParserAjax)|(1<<BeetlParserFragment)|(1<<BeetlParserLEFT_BRACE)|(1<<BeetlParserLEFT_PAR)|(1<<BeetlParserLEFT_SQBR)|(1<<BeetlParserINCREASE)|(1<<BeetlParserDECREASE)|(1<<BeetlParserADD)|(1<<BeetlParserMIN))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(BeetlParserNOT-40))|(1<<(BeetlParserAT-40))|(1<<(BeetlParserNULL-40))|(1<<(BeetlParserTRUE-40))|(1<<(BeetlParserFALSE-40))|(1<<(BeetlParserEND-40))|(1<<(BeetlParserLEFT_TOKEN-40))|(1<<(BeetlParserLEFT_TOKEN2-40))|(1<<(BeetlParserLEFT_TEXT_TOKEN-40))|(1<<(BeetlParserDecimalLiteral-40))|(1<<(BeetlParserFloatingPointLiteral-40))|(1<<(BeetlParserStringLiteral-40))|(1<<(BeetlParserIdentifier-40))|(1<<(BeetlParserCOMMENT_TAG-40)))) != 0) {
		{
			p.SetState(100)
			p.Statement()
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(106)
		p.Match(BeetlParserEOF)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LEFT_BRACE() antlr.TerminalNode {
	return s.GetToken(BeetlParserLEFT_BRACE, 0)
}

func (s *BlockContext) RIGHT_BRACE() antlr.TerminalNode {
	return s.GetToken(BeetlParserRIGHT_BRACE, 0)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *BeetlParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BeetlParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(BeetlParserLEFT_BRACE)
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeetlParserIf)|(1<<BeetlParserFor)|(1<<BeetlParserWhile)|(1<<BeetlParserSwitch)|(1<<BeetlParserSelect)|(1<<BeetlParserReturn)|(1<<BeetlParserBreak)|(1<<BeetlParserContinue)|(1<<BeetlParserVar)|(1<<BeetlParserDirective)|(1<<BeetlParserTry)|(1<<BeetlParserAjax)|(1<<BeetlParserFragment)|(1<<BeetlParserLEFT_BRACE)|(1<<BeetlParserLEFT_PAR)|(1<<BeetlParserLEFT_SQBR)|(1<<BeetlParserINCREASE)|(1<<BeetlParserDECREASE)|(1<<BeetlParserADD)|(1<<BeetlParserMIN))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(BeetlParserNOT-40))|(1<<(BeetlParserAT-40))|(1<<(BeetlParserNULL-40))|(1<<(BeetlParserTRUE-40))|(1<<(BeetlParserFALSE-40))|(1<<(BeetlParserEND-40))|(1<<(BeetlParserLEFT_TOKEN-40))|(1<<(BeetlParserLEFT_TOKEN2-40))|(1<<(BeetlParserLEFT_TEXT_TOKEN-40))|(1<<(BeetlParserDecimalLiteral-40))|(1<<(BeetlParserFloatingPointLiteral-40))|(1<<(BeetlParserStringLiteral-40))|(1<<(BeetlParserIdentifier-40))|(1<<(BeetlParserCOMMENT_TAG-40)))) != 0) {
		{
			p.SetState(109)
			p.Statement()
		}

		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(115)
		p.Match(BeetlParserRIGHT_BRACE)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyFrom(ctx *StatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StaticOutputStContext struct {
	*StatementContext
}

func NewStaticOutputStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StaticOutputStContext {
	var p = new(StaticOutputStContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *StaticOutputStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StaticOutputStContext) ConstantsTextStatment() IConstantsTextStatmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantsTextStatmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantsTextStatmentContext)
}

func (s *StaticOutputStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterStaticOutputSt(s)
	}
}

func (s *StaticOutputStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitStaticOutputSt(s)
	}
}

type ReturnStContext struct {
	*StatementContext
}

func NewReturnStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStContext {
	var p = new(ReturnStContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ReturnStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStContext) Return() antlr.TerminalNode {
	return s.GetToken(BeetlParserReturn, 0)
}

func (s *ReturnStContext) END() antlr.TerminalNode {
	return s.GetToken(BeetlParserEND, 0)
}

func (s *ReturnStContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterReturnSt(s)
	}
}

func (s *ReturnStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitReturnSt(s)
	}
}

type TryStContext struct {
	*StatementContext
}

func NewTryStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TryStContext {
	var p = new(TryStContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *TryStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TryStContext) Try() antlr.TerminalNode {
	return s.GetToken(BeetlParserTry, 0)
}

func (s *TryStContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *TryStContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *TryStContext) Catch() antlr.TerminalNode {
	return s.GetToken(BeetlParserCatch, 0)
}

func (s *TryStContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(BeetlParserLEFT_PAR, 0)
}

func (s *TryStContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(BeetlParserRIGHT_PAR, 0)
}

func (s *TryStContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BeetlParserIdentifier, 0)
}

func (s *TryStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterTrySt(s)
	}
}

func (s *TryStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitTrySt(s)
	}
}

type VarStContext struct {
	*StatementContext
}

func NewVarStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarStContext {
	var p = new(VarStContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *VarStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarStContext) Var() antlr.TerminalNode {
	return s.GetToken(BeetlParserVar, 0)
}

func (s *VarStContext) VarDeclareList() IVarDeclareListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDeclareListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarDeclareListContext)
}

func (s *VarStContext) END() antlr.TerminalNode {
	return s.GetToken(BeetlParserEND, 0)
}

func (s *VarStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterVarSt(s)
	}
}

func (s *VarStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitVarSt(s)
	}
}

type AssignStContext struct {
	*StatementContext
}

func NewAssignStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignStContext {
	var p = new(AssignStContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *AssignStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStContext) AssignMent() IAssignMentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignMentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignMentContext)
}

func (s *AssignStContext) END() antlr.TerminalNode {
	return s.GetToken(BeetlParserEND, 0)
}

func (s *AssignStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterAssignSt(s)
	}
}

func (s *AssignStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitAssignSt(s)
	}
}

type WhileStContext struct {
	*StatementContext
}

func NewWhileStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileStContext {
	var p = new(WhileStContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *WhileStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStContext) While() antlr.TerminalNode {
	return s.GetToken(BeetlParserWhile, 0)
}

func (s *WhileStContext) ParExpression() IParExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParExpressionContext)
}

func (s *WhileStContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *WhileStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterWhileSt(s)
	}
}

func (s *WhileStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitWhileSt(s)
	}
}

type FunctionTagStContext struct {
	*StatementContext
}

func NewFunctionTagStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionTagStContext {
	var p = new(FunctionTagStContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *FunctionTagStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionTagStContext) FunctionTagCall() IFunctionTagCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionTagCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionTagCallContext)
}

func (s *FunctionTagStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterFunctionTagSt(s)
	}
}

func (s *FunctionTagStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitFunctionTagSt(s)
	}
}

type BreakStContext struct {
	*StatementContext
}

func NewBreakStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStContext {
	var p = new(BreakStContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *BreakStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStContext) Break() antlr.TerminalNode {
	return s.GetToken(BeetlParserBreak, 0)
}

func (s *BreakStContext) END() antlr.TerminalNode {
	return s.GetToken(BeetlParserEND, 0)
}

func (s *BreakStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterBreakSt(s)
	}
}

func (s *BreakStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitBreakSt(s)
	}
}

type AjaxStContext struct {
	*StatementContext
}

func NewAjaxStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AjaxStContext {
	var p = new(AjaxStContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *AjaxStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AjaxStContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(BeetlParserIdentifier)
}

func (s *AjaxStContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(BeetlParserIdentifier, i)
}

func (s *AjaxStContext) COLON() antlr.TerminalNode {
	return s.GetToken(BeetlParserCOLON, 0)
}

func (s *AjaxStContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *AjaxStContext) Ajax() antlr.TerminalNode {
	return s.GetToken(BeetlParserAjax, 0)
}

func (s *AjaxStContext) Fragment() antlr.TerminalNode {
	return s.GetToken(BeetlParserFragment, 0)
}

func (s *AjaxStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterAjaxSt(s)
	}
}

func (s *AjaxStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitAjaxSt(s)
	}
}

type StatmentExpStContext struct {
	*StatementContext
}

func NewStatmentExpStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatmentExpStContext {
	var p = new(StatmentExpStContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *StatmentExpStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatmentExpStContext) StatementExpression() IStatementExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementExpressionContext)
}

func (s *StatmentExpStContext) END() antlr.TerminalNode {
	return s.GetToken(BeetlParserEND, 0)
}

func (s *StatmentExpStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterStatmentExpSt(s)
	}
}

func (s *StatmentExpStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitStatmentExpSt(s)
	}
}

type IfStContext struct {
	*StatementContext
}

func NewIfStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStContext {
	var p = new(IfStContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *IfStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStContext) If() antlr.TerminalNode {
	return s.GetToken(BeetlParserIf, 0)
}

func (s *IfStContext) ParExpression() IParExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParExpressionContext)
}

func (s *IfStContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *IfStContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfStContext) Else() antlr.TerminalNode {
	return s.GetToken(BeetlParserElse, 0)
}

func (s *IfStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterIfSt(s)
	}
}

func (s *IfStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitIfSt(s)
	}
}

type DirectiveStContext struct {
	*StatementContext
}

func NewDirectiveStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DirectiveStContext {
	var p = new(DirectiveStContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *DirectiveStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveStContext) Directive() antlr.TerminalNode {
	return s.GetToken(BeetlParserDirective, 0)
}

func (s *DirectiveStContext) DirectiveExp() IDirectiveExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveExpContext)
}

func (s *DirectiveStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterDirectiveSt(s)
	}
}

func (s *DirectiveStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitDirectiveSt(s)
	}
}

type ContinueStContext struct {
	*StatementContext
}

func NewContinueStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStContext {
	var p = new(ContinueStContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ContinueStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStContext) Continue() antlr.TerminalNode {
	return s.GetToken(BeetlParserContinue, 0)
}

func (s *ContinueStContext) END() antlr.TerminalNode {
	return s.GetToken(BeetlParserEND, 0)
}

func (s *ContinueStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterContinueSt(s)
	}
}

func (s *ContinueStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitContinueSt(s)
	}
}

type SiwchStContext struct {
	*StatementContext
}

func NewSiwchStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SiwchStContext {
	var p = new(SiwchStContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *SiwchStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SiwchStContext) Switch() antlr.TerminalNode {
	return s.GetToken(BeetlParserSwitch, 0)
}

func (s *SiwchStContext) ParExpression() IParExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParExpressionContext)
}

func (s *SiwchStContext) SwitchBlock() ISwitchBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitchBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISwitchBlockContext)
}

func (s *SiwchStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterSiwchSt(s)
	}
}

func (s *SiwchStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitSiwchSt(s)
	}
}

type CommentTagStContext struct {
	*StatementContext
}

func NewCommentTagStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CommentTagStContext {
	var p = new(CommentTagStContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *CommentTagStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentTagStContext) COMMENT_TAG() antlr.TerminalNode {
	return s.GetToken(BeetlParserCOMMENT_TAG, 0)
}

func (s *CommentTagStContext) CommentTypeTag() ICommentTypeTagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentTypeTagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentTypeTagContext)
}

func (s *CommentTagStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterCommentTagSt(s)
	}
}

func (s *CommentTagStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitCommentTagSt(s)
	}
}

type BlockStContext struct {
	*StatementContext
}

func NewBlockStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockStContext {
	var p = new(BlockStContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *BlockStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *BlockStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterBlockSt(s)
	}
}

func (s *BlockStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitBlockSt(s)
	}
}

type SelectStContext struct {
	*StatementContext
}

func NewSelectStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectStContext {
	var p = new(SelectStContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *SelectStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectStContext) Select() antlr.TerminalNode {
	return s.GetToken(BeetlParserSelect, 0)
}

func (s *SelectStContext) G_switchStatment() IG_switchStatmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IG_switchStatmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IG_switchStatmentContext)
}

func (s *SelectStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterSelectSt(s)
	}
}

func (s *SelectStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitSelectSt(s)
	}
}

type EndContext struct {
	*StatementContext
}

func NewEndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EndContext {
	var p = new(EndContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *EndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndContext) END() antlr.TerminalNode {
	return s.GetToken(BeetlParserEND, 0)
}

func (s *EndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterEnd(s)
	}
}

func (s *EndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitEnd(s)
	}
}

type ForStContext struct {
	*StatementContext
}

func NewForStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForStContext {
	var p = new(ForStContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ForStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStContext) For() antlr.TerminalNode {
	return s.GetToken(BeetlParserFor, 0)
}

func (s *ForStContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(BeetlParserLEFT_PAR, 0)
}

func (s *ForStContext) ForControl() IForControlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForControlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForControlContext)
}

func (s *ForStContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(BeetlParserRIGHT_PAR, 0)
}

func (s *ForStContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ForStContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ForStContext) Elsefor() antlr.TerminalNode {
	return s.GetToken(BeetlParserElsefor, 0)
}

func (s *ForStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterForSt(s)
	}
}

func (s *ForStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitForSt(s)
	}
}

type TextOutputStContext struct {
	*StatementContext
}

func NewTextOutputStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TextOutputStContext {
	var p = new(TextOutputStContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *TextOutputStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextOutputStContext) TextStatment() ITextStatmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextStatmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITextStatmentContext)
}

func (s *TextOutputStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterTextOutputSt(s)
	}
}

func (s *TextOutputStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitTextOutputSt(s)
	}
}

func (p *BeetlParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BeetlParserRULE_statement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBlockStContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(117)
			p.Block()
		}

	case 2:
		localctx = NewTextOutputStContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.TextStatment()
		}

	case 3:
		localctx = NewStaticOutputStContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(119)
			p.ConstantsTextStatment()
		}

	case 4:
		localctx = NewCommentTagStContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(120)
			p.Match(BeetlParserCOMMENT_TAG)
		}
		{
			p.SetState(121)
			p.CommentTypeTag()
		}

	case 5:
		localctx = NewIfStContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(122)
			p.Match(BeetlParserIf)
		}
		{
			p.SetState(123)
			p.ParExpression()
		}
		{
			p.SetState(124)
			p.Statement()
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(125)
				p.Match(BeetlParserElse)
			}
			{
				p.SetState(126)
				p.Statement()
			}

		}

	case 6:
		localctx = NewForStContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(129)
			p.Match(BeetlParserFor)
		}
		{
			p.SetState(130)
			p.Match(BeetlParserLEFT_PAR)
		}
		{
			p.SetState(131)
			p.ForControl()
		}
		{
			p.SetState(132)
			p.Match(BeetlParserRIGHT_PAR)
		}
		{
			p.SetState(133)
			p.Statement()
		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(134)
				p.Match(BeetlParserElsefor)
			}
			{
				p.SetState(135)
				p.Statement()
			}

		}

	case 7:
		localctx = NewWhileStContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(138)
			p.Match(BeetlParserWhile)
		}
		{
			p.SetState(139)
			p.ParExpression()
		}
		{
			p.SetState(140)
			p.Statement()
		}

	case 8:
		localctx = NewSiwchStContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(142)
			p.Match(BeetlParserSwitch)
		}
		{
			p.SetState(143)
			p.ParExpression()
		}
		{
			p.SetState(144)
			p.SwitchBlock()
		}

	case 9:
		localctx = NewSelectStContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(146)
			p.Match(BeetlParserSelect)
		}
		{
			p.SetState(147)
			p.G_switchStatment()
		}

	case 10:
		localctx = NewTryStContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(148)
			p.Match(BeetlParserTry)
		}
		{
			p.SetState(149)
			p.Block()
		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BeetlParserCatch {
			{
				p.SetState(150)
				p.Match(BeetlParserCatch)
			}
			{
				p.SetState(151)
				p.Match(BeetlParserLEFT_PAR)
			}
			p.SetState(153)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == BeetlParserIdentifier {
				{
					p.SetState(152)
					p.Match(BeetlParserIdentifier)
				}

			}
			{
				p.SetState(155)
				p.Match(BeetlParserRIGHT_PAR)
			}
			{
				p.SetState(156)
				p.Block()
			}

		}

	case 11:
		localctx = NewReturnStContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(159)
			p.Match(BeetlParserReturn)
		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeetlParserLEFT_BRACE)|(1<<BeetlParserLEFT_PAR)|(1<<BeetlParserLEFT_SQBR)|(1<<BeetlParserINCREASE)|(1<<BeetlParserDECREASE)|(1<<BeetlParserADD)|(1<<BeetlParserMIN))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(BeetlParserNOT-40))|(1<<(BeetlParserAT-40))|(1<<(BeetlParserNULL-40))|(1<<(BeetlParserTRUE-40))|(1<<(BeetlParserFALSE-40))|(1<<(BeetlParserDecimalLiteral-40))|(1<<(BeetlParserFloatingPointLiteral-40))|(1<<(BeetlParserStringLiteral-40))|(1<<(BeetlParserIdentifier-40)))) != 0) {
			{
				p.SetState(160)
				p.expression(0)
			}

		}
		{
			p.SetState(163)
			p.Match(BeetlParserEND)
		}

	case 12:
		localctx = NewBreakStContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(164)
			p.Match(BeetlParserBreak)
		}
		{
			p.SetState(165)
			p.Match(BeetlParserEND)
		}

	case 13:
		localctx = NewContinueStContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(166)
			p.Match(BeetlParserContinue)
		}
		{
			p.SetState(167)
			p.Match(BeetlParserEND)
		}

	case 14:
		localctx = NewVarStContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(168)
			p.Match(BeetlParserVar)
		}
		{
			p.SetState(169)
			p.VarDeclareList()
		}
		{
			p.SetState(170)
			p.Match(BeetlParserEND)
		}

	case 15:
		localctx = NewDirectiveStContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(172)
			p.Match(BeetlParserDirective)
		}
		{
			p.SetState(173)
			p.DirectiveExp()
		}

	case 16:
		localctx = NewAssignStContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(174)
			p.AssignMent()
		}
		{
			p.SetState(175)
			p.Match(BeetlParserEND)
		}

	case 17:
		localctx = NewFunctionTagStContext(p, localctx)
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(177)
			p.FunctionTagCall()
		}

	case 18:
		localctx = NewStatmentExpStContext(p, localctx)
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(178)
			p.StatementExpression()
		}
		{
			p.SetState(179)
			p.Match(BeetlParserEND)
		}

	case 19:
		localctx = NewAjaxStContext(p, localctx)
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(181)
			_la = p.GetTokenStream().LA(1)

			if !(_la == BeetlParserAjax || _la == BeetlParserFragment) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(182)
			p.Match(BeetlParserIdentifier)
		}
		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BeetlParserIdentifier {
			{
				p.SetState(183)
				p.Match(BeetlParserIdentifier)
			}

		}
		{
			p.SetState(186)
			p.Match(BeetlParserCOLON)
		}
		{
			p.SetState(187)
			p.Block()
		}

	case 20:
		localctx = NewEndContext(p, localctx)
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(188)
			p.Match(BeetlParserEND)
		}

	}

	return localctx
}

// ICommentTypeTagContext is an interface to support dynamic dispatch.
type ICommentTypeTagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentTypeTagContext differentiates from other interfaces.
	IsCommentTypeTagContext()
}

type CommentTypeTagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentTypeTagContext() *CommentTypeTagContext {
	var p = new(CommentTypeTagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_commentTypeTag
	return p
}

func (*CommentTypeTagContext) IsCommentTypeTagContext() {}

func NewCommentTypeTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentTypeTagContext {
	var p = new(CommentTypeTagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_commentTypeTag

	return p
}

func (s *CommentTypeTagContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentTypeTagContext) LEFT_PAR1() antlr.TerminalNode {
	return s.GetToken(BeetlParserLEFT_PAR1, 0)
}

func (s *CommentTypeTagContext) AllCommentTypeItemTag() []ICommentTypeItemTagContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentTypeItemTagContext)(nil)).Elem())
	var tst = make([]ICommentTypeItemTagContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentTypeItemTagContext)
		}
	}

	return tst
}

func (s *CommentTypeTagContext) CommentTypeItemTag(i int) ICommentTypeItemTagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentTypeItemTagContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentTypeItemTagContext)
}

func (s *CommentTypeTagContext) RIGHT_PAR1() antlr.TerminalNode {
	return s.GetToken(BeetlParserRIGHT_PAR1, 0)
}

func (s *CommentTypeTagContext) AllCOMMA1() []antlr.TerminalNode {
	return s.GetTokens(BeetlParserCOMMA1)
}

func (s *CommentTypeTagContext) COMMA1(i int) antlr.TerminalNode {
	return s.GetToken(BeetlParserCOMMA1, i)
}

func (s *CommentTypeTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentTypeTagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentTypeTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterCommentTypeTag(s)
	}
}

func (s *CommentTypeTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitCommentTypeTag(s)
	}
}

func (p *BeetlParser) CommentTypeTag() (localctx ICommentTypeTagContext) {
	localctx = NewCommentTypeTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BeetlParserRULE_commentTypeTag)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.Match(BeetlParserLEFT_PAR1)
	}
	{
		p.SetState(192)
		p.CommentTypeItemTag()
	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BeetlParserCOMMA1 {
		{
			p.SetState(193)
			p.Match(BeetlParserCOMMA1)
		}
		{
			p.SetState(194)
			p.CommentTypeItemTag()
		}

		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(200)
		p.Match(BeetlParserRIGHT_PAR1)
	}

	return localctx
}

// ICommentTypeItemTagContext is an interface to support dynamic dispatch.
type ICommentTypeItemTagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentTypeItemTagContext differentiates from other interfaces.
	IsCommentTypeItemTagContext()
}

type CommentTypeItemTagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentTypeItemTagContext() *CommentTypeItemTagContext {
	var p = new(CommentTypeItemTagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_commentTypeItemTag
	return p
}

func (*CommentTypeItemTagContext) IsCommentTypeItemTagContext() {}

func NewCommentTypeItemTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentTypeItemTagContext {
	var p = new(CommentTypeItemTagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_commentTypeItemTag

	return p
}

func (s *CommentTypeItemTagContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentTypeItemTagContext) ClassOrInterfaceType() IClassOrInterfaceTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassOrInterfaceTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassOrInterfaceTypeContext)
}

func (s *CommentTypeItemTagContext) Identifier1() antlr.TerminalNode {
	return s.GetToken(BeetlParserIdentifier1, 0)
}

func (s *CommentTypeItemTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentTypeItemTagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentTypeItemTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterCommentTypeItemTag(s)
	}
}

func (s *CommentTypeItemTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitCommentTypeItemTag(s)
	}
}

func (p *BeetlParser) CommentTypeItemTag() (localctx ICommentTypeItemTagContext) {
	localctx = NewCommentTypeItemTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BeetlParserRULE_commentTypeItemTag)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.ClassOrInterfaceType()
	}
	{
		p.SetState(203)
		p.Match(BeetlParserIdentifier1)
	}

	return localctx
}

// IClassOrInterfaceTypeContext is an interface to support dynamic dispatch.
type IClassOrInterfaceTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassOrInterfaceTypeContext differentiates from other interfaces.
	IsClassOrInterfaceTypeContext()
}

type ClassOrInterfaceTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassOrInterfaceTypeContext() *ClassOrInterfaceTypeContext {
	var p = new(ClassOrInterfaceTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_classOrInterfaceType
	return p
}

func (*ClassOrInterfaceTypeContext) IsClassOrInterfaceTypeContext() {}

func NewClassOrInterfaceTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassOrInterfaceTypeContext {
	var p = new(ClassOrInterfaceTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_classOrInterfaceType

	return p
}

func (s *ClassOrInterfaceTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassOrInterfaceTypeContext) AllIdentifier1() []antlr.TerminalNode {
	return s.GetTokens(BeetlParserIdentifier1)
}

func (s *ClassOrInterfaceTypeContext) Identifier1(i int) antlr.TerminalNode {
	return s.GetToken(BeetlParserIdentifier1, i)
}

func (s *ClassOrInterfaceTypeContext) AllPERIOD1() []antlr.TerminalNode {
	return s.GetTokens(BeetlParserPERIOD1)
}

func (s *ClassOrInterfaceTypeContext) PERIOD1(i int) antlr.TerminalNode {
	return s.GetToken(BeetlParserPERIOD1, i)
}

func (s *ClassOrInterfaceTypeContext) TypeArguments() ITypeArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeArgumentsContext)
}

func (s *ClassOrInterfaceTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassOrInterfaceTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassOrInterfaceTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterClassOrInterfaceType(s)
	}
}

func (s *ClassOrInterfaceTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitClassOrInterfaceType(s)
	}
}

func (p *BeetlParser) ClassOrInterfaceType() (localctx IClassOrInterfaceTypeContext) {
	localctx = NewClassOrInterfaceTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BeetlParserRULE_classOrInterfaceType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(BeetlParserIdentifier1)
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BeetlParserPERIOD1 {
		{
			p.SetState(206)
			p.Match(BeetlParserPERIOD1)
		}
		{
			p.SetState(207)
			p.Match(BeetlParserIdentifier1)
		}

		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BeetlParserLEFT_ANGULAR {
		{
			p.SetState(213)
			p.TypeArguments()
		}

	}

	return localctx
}

// ITypeArgumentsContext is an interface to support dynamic dispatch.
type ITypeArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeArgumentsContext differentiates from other interfaces.
	IsTypeArgumentsContext()
}

type TypeArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeArgumentsContext() *TypeArgumentsContext {
	var p = new(TypeArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_typeArguments
	return p
}

func (*TypeArgumentsContext) IsTypeArgumentsContext() {}

func NewTypeArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeArgumentsContext {
	var p = new(TypeArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_typeArguments

	return p
}

func (s *TypeArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeArgumentsContext) LEFT_ANGULAR() antlr.TerminalNode {
	return s.GetToken(BeetlParserLEFT_ANGULAR, 0)
}

func (s *TypeArgumentsContext) AllTypeArgument() []ITypeArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeArgumentContext)(nil)).Elem())
	var tst = make([]ITypeArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeArgumentContext)
		}
	}

	return tst
}

func (s *TypeArgumentsContext) TypeArgument(i int) ITypeArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeArgumentContext)
}

func (s *TypeArgumentsContext) RIGHT_ANGULAR() antlr.TerminalNode {
	return s.GetToken(BeetlParserRIGHT_ANGULAR, 0)
}

func (s *TypeArgumentsContext) AllCOMMA1() []antlr.TerminalNode {
	return s.GetTokens(BeetlParserCOMMA1)
}

func (s *TypeArgumentsContext) COMMA1(i int) antlr.TerminalNode {
	return s.GetToken(BeetlParserCOMMA1, i)
}

func (s *TypeArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterTypeArguments(s)
	}
}

func (s *TypeArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitTypeArguments(s)
	}
}

func (p *BeetlParser) TypeArguments() (localctx ITypeArgumentsContext) {
	localctx = NewTypeArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BeetlParserRULE_typeArguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Match(BeetlParserLEFT_ANGULAR)
	}
	{
		p.SetState(217)
		p.TypeArgument()
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BeetlParserCOMMA1 {
		{
			p.SetState(218)
			p.Match(BeetlParserCOMMA1)
		}
		{
			p.SetState(219)
			p.TypeArgument()
		}

		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(225)
		p.Match(BeetlParserRIGHT_ANGULAR)
	}

	return localctx
}

// ITypeArgumentContext is an interface to support dynamic dispatch.
type ITypeArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeArgumentContext differentiates from other interfaces.
	IsTypeArgumentContext()
}

type TypeArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeArgumentContext() *TypeArgumentContext {
	var p = new(TypeArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_typeArgument
	return p
}

func (*TypeArgumentContext) IsTypeArgumentContext() {}

func NewTypeArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeArgumentContext {
	var p = new(TypeArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_typeArgument

	return p
}

func (s *TypeArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeArgumentContext) ClassOrInterfaceType() IClassOrInterfaceTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassOrInterfaceTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassOrInterfaceTypeContext)
}

func (s *TypeArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterTypeArgument(s)
	}
}

func (s *TypeArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitTypeArgument(s)
	}
}

func (p *BeetlParser) TypeArgument() (localctx ITypeArgumentContext) {
	localctx = NewTypeArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BeetlParserRULE_typeArgument)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.ClassOrInterfaceType()
	}

	return localctx
}

// IDirectiveExpContext is an interface to support dynamic dispatch.
type IDirectiveExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveExpContext differentiates from other interfaces.
	IsDirectiveExpContext()
}

type DirectiveExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveExpContext() *DirectiveExpContext {
	var p = new(DirectiveExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_directiveExp
	return p
}

func (*DirectiveExpContext) IsDirectiveExpContext() {}

func NewDirectiveExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveExpContext {
	var p = new(DirectiveExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_directiveExp

	return p
}

func (s *DirectiveExpContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveExpContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BeetlParserIdentifier, 0)
}

func (s *DirectiveExpContext) END() antlr.TerminalNode {
	return s.GetToken(BeetlParserEND, 0)
}

func (s *DirectiveExpContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(BeetlParserStringLiteral, 0)
}

func (s *DirectiveExpContext) DirectiveExpIDList() IDirectiveExpIDListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveExpIDListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveExpIDListContext)
}

func (s *DirectiveExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterDirectiveExp(s)
	}
}

func (s *DirectiveExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitDirectiveExp(s)
	}
}

func (p *BeetlParser) DirectiveExp() (localctx IDirectiveExpContext) {
	localctx = NewDirectiveExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BeetlParserRULE_directiveExp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.Match(BeetlParserIdentifier)
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeetlParserStringLiteral:
		{
			p.SetState(230)
			p.Match(BeetlParserStringLiteral)
		}

	case BeetlParserIdentifier:
		{
			p.SetState(231)
			p.DirectiveExpIDList()
		}

	case BeetlParserEND:

	default:
	}
	{
		p.SetState(234)
		p.Match(BeetlParserEND)
	}

	return localctx
}

// IDirectiveExpIDListContext is an interface to support dynamic dispatch.
type IDirectiveExpIDListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveExpIDListContext differentiates from other interfaces.
	IsDirectiveExpIDListContext()
}

type DirectiveExpIDListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveExpIDListContext() *DirectiveExpIDListContext {
	var p = new(DirectiveExpIDListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_directiveExpIDList
	return p
}

func (*DirectiveExpIDListContext) IsDirectiveExpIDListContext() {}

func NewDirectiveExpIDListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveExpIDListContext {
	var p = new(DirectiveExpIDListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_directiveExpIDList

	return p
}

func (s *DirectiveExpIDListContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveExpIDListContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(BeetlParserIdentifier)
}

func (s *DirectiveExpIDListContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(BeetlParserIdentifier, i)
}

func (s *DirectiveExpIDListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(BeetlParserCOMMA)
}

func (s *DirectiveExpIDListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(BeetlParserCOMMA, i)
}

func (s *DirectiveExpIDListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveExpIDListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveExpIDListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterDirectiveExpIDList(s)
	}
}

func (s *DirectiveExpIDListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitDirectiveExpIDList(s)
	}
}

func (p *BeetlParser) DirectiveExpIDList() (localctx IDirectiveExpIDListContext) {
	localctx = NewDirectiveExpIDListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BeetlParserRULE_directiveExpIDList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.Match(BeetlParserIdentifier)
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BeetlParserCOMMA {
		{
			p.SetState(237)
			p.Match(BeetlParserCOMMA)
		}
		{
			p.SetState(238)
			p.Match(BeetlParserIdentifier)
		}

		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IG_switchStatmentContext is an interface to support dynamic dispatch.
type IG_switchStatmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBase returns the base rule contexts.
	GetBase() IExpressionContext

	// SetBase sets the base rule contexts.
	SetBase(IExpressionContext)

	// IsG_switchStatmentContext differentiates from other interfaces.
	IsG_switchStatmentContext()
}

type G_switchStatmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	base   IExpressionContext
}

func NewEmptyG_switchStatmentContext() *G_switchStatmentContext {
	var p = new(G_switchStatmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_g_switchStatment
	return p
}

func (*G_switchStatmentContext) IsG_switchStatmentContext() {}

func NewG_switchStatmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *G_switchStatmentContext {
	var p = new(G_switchStatmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_g_switchStatment

	return p
}

func (s *G_switchStatmentContext) GetParser() antlr.Parser { return s.parser }

func (s *G_switchStatmentContext) GetBase() IExpressionContext { return s.base }

func (s *G_switchStatmentContext) SetBase(v IExpressionContext) { s.base = v }

func (s *G_switchStatmentContext) LEFT_BRACE() antlr.TerminalNode {
	return s.GetToken(BeetlParserLEFT_BRACE, 0)
}

func (s *G_switchStatmentContext) RIGHT_BRACE() antlr.TerminalNode {
	return s.GetToken(BeetlParserRIGHT_BRACE, 0)
}

func (s *G_switchStatmentContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(BeetlParserLEFT_PAR, 0)
}

func (s *G_switchStatmentContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(BeetlParserRIGHT_PAR, 0)
}

func (s *G_switchStatmentContext) AllG_caseStatment() []IG_caseStatmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IG_caseStatmentContext)(nil)).Elem())
	var tst = make([]IG_caseStatmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IG_caseStatmentContext)
		}
	}

	return tst
}

func (s *G_switchStatmentContext) G_caseStatment(i int) IG_caseStatmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IG_caseStatmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IG_caseStatmentContext)
}

func (s *G_switchStatmentContext) G_defaultStatment() IG_defaultStatmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IG_defaultStatmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IG_defaultStatmentContext)
}

func (s *G_switchStatmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *G_switchStatmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *G_switchStatmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *G_switchStatmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterG_switchStatment(s)
	}
}

func (s *G_switchStatmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitG_switchStatment(s)
	}
}

func (p *BeetlParser) G_switchStatment() (localctx IG_switchStatmentContext) {
	localctx = NewG_switchStatmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BeetlParserRULE_g_switchStatment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BeetlParserLEFT_PAR {
		{
			p.SetState(244)
			p.Match(BeetlParserLEFT_PAR)
		}
		{
			p.SetState(245)

			var _x = p.expression(0)

			localctx.(*G_switchStatmentContext).base = _x
		}
		{
			p.SetState(246)
			p.Match(BeetlParserRIGHT_PAR)
		}

	}
	{
		p.SetState(250)
		p.Match(BeetlParserLEFT_BRACE)
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BeetlParserCase {
		{
			p.SetState(251)
			p.G_caseStatment()
		}

		p.SetState(256)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BeetlParserDefault {
		{
			p.SetState(257)
			p.G_defaultStatment()
		}

	}
	{
		p.SetState(260)
		p.Match(BeetlParserRIGHT_BRACE)
	}

	return localctx
}

// IG_caseStatmentContext is an interface to support dynamic dispatch.
type IG_caseStatmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsG_caseStatmentContext differentiates from other interfaces.
	IsG_caseStatmentContext()
}

type G_caseStatmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyG_caseStatmentContext() *G_caseStatmentContext {
	var p = new(G_caseStatmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_g_caseStatment
	return p
}

func (*G_caseStatmentContext) IsG_caseStatmentContext() {}

func NewG_caseStatmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *G_caseStatmentContext {
	var p = new(G_caseStatmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_g_caseStatment

	return p
}

func (s *G_caseStatmentContext) GetParser() antlr.Parser { return s.parser }

func (s *G_caseStatmentContext) Case() antlr.TerminalNode {
	return s.GetToken(BeetlParserCase, 0)
}

func (s *G_caseStatmentContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *G_caseStatmentContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *G_caseStatmentContext) COLON() antlr.TerminalNode {
	return s.GetToken(BeetlParserCOLON, 0)
}

func (s *G_caseStatmentContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(BeetlParserCOMMA)
}

func (s *G_caseStatmentContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(BeetlParserCOMMA, i)
}

func (s *G_caseStatmentContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *G_caseStatmentContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *G_caseStatmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *G_caseStatmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *G_caseStatmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterG_caseStatment(s)
	}
}

func (s *G_caseStatmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitG_caseStatment(s)
	}
}

func (p *BeetlParser) G_caseStatment() (localctx IG_caseStatmentContext) {
	localctx = NewG_caseStatmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BeetlParserRULE_g_caseStatment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
		p.Match(BeetlParserCase)
	}
	{
		p.SetState(263)
		p.expression(0)
	}
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BeetlParserCOMMA {
		{
			p.SetState(264)
			p.Match(BeetlParserCOMMA)
		}
		{
			p.SetState(265)
			p.expression(0)
		}

		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(271)
		p.Match(BeetlParserCOLON)
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeetlParserIf)|(1<<BeetlParserFor)|(1<<BeetlParserWhile)|(1<<BeetlParserSwitch)|(1<<BeetlParserSelect)|(1<<BeetlParserReturn)|(1<<BeetlParserBreak)|(1<<BeetlParserContinue)|(1<<BeetlParserVar)|(1<<BeetlParserDirective)|(1<<BeetlParserTry)|(1<<BeetlParserAjax)|(1<<BeetlParserFragment)|(1<<BeetlParserLEFT_BRACE)|(1<<BeetlParserLEFT_PAR)|(1<<BeetlParserLEFT_SQBR)|(1<<BeetlParserINCREASE)|(1<<BeetlParserDECREASE)|(1<<BeetlParserADD)|(1<<BeetlParserMIN))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(BeetlParserNOT-40))|(1<<(BeetlParserAT-40))|(1<<(BeetlParserNULL-40))|(1<<(BeetlParserTRUE-40))|(1<<(BeetlParserFALSE-40))|(1<<(BeetlParserEND-40))|(1<<(BeetlParserLEFT_TOKEN-40))|(1<<(BeetlParserLEFT_TOKEN2-40))|(1<<(BeetlParserLEFT_TEXT_TOKEN-40))|(1<<(BeetlParserDecimalLiteral-40))|(1<<(BeetlParserFloatingPointLiteral-40))|(1<<(BeetlParserStringLiteral-40))|(1<<(BeetlParserIdentifier-40))|(1<<(BeetlParserCOMMENT_TAG-40)))) != 0) {
		{
			p.SetState(272)
			p.Statement()
		}

		p.SetState(277)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IG_defaultStatmentContext is an interface to support dynamic dispatch.
type IG_defaultStatmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsG_defaultStatmentContext differentiates from other interfaces.
	IsG_defaultStatmentContext()
}

type G_defaultStatmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyG_defaultStatmentContext() *G_defaultStatmentContext {
	var p = new(G_defaultStatmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_g_defaultStatment
	return p
}

func (*G_defaultStatmentContext) IsG_defaultStatmentContext() {}

func NewG_defaultStatmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *G_defaultStatmentContext {
	var p = new(G_defaultStatmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_g_defaultStatment

	return p
}

func (s *G_defaultStatmentContext) GetParser() antlr.Parser { return s.parser }

func (s *G_defaultStatmentContext) Default() antlr.TerminalNode {
	return s.GetToken(BeetlParserDefault, 0)
}

func (s *G_defaultStatmentContext) COLON() antlr.TerminalNode {
	return s.GetToken(BeetlParserCOLON, 0)
}

func (s *G_defaultStatmentContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *G_defaultStatmentContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *G_defaultStatmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *G_defaultStatmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *G_defaultStatmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterG_defaultStatment(s)
	}
}

func (s *G_defaultStatmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitG_defaultStatment(s)
	}
}

func (p *BeetlParser) G_defaultStatment() (localctx IG_defaultStatmentContext) {
	localctx = NewG_defaultStatmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BeetlParserRULE_g_defaultStatment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(278)
		p.Match(BeetlParserDefault)
	}
	{
		p.SetState(279)
		p.Match(BeetlParserCOLON)
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeetlParserIf)|(1<<BeetlParserFor)|(1<<BeetlParserWhile)|(1<<BeetlParserSwitch)|(1<<BeetlParserSelect)|(1<<BeetlParserReturn)|(1<<BeetlParserBreak)|(1<<BeetlParserContinue)|(1<<BeetlParserVar)|(1<<BeetlParserDirective)|(1<<BeetlParserTry)|(1<<BeetlParserAjax)|(1<<BeetlParserFragment)|(1<<BeetlParserLEFT_BRACE)|(1<<BeetlParserLEFT_PAR)|(1<<BeetlParserLEFT_SQBR)|(1<<BeetlParserINCREASE)|(1<<BeetlParserDECREASE)|(1<<BeetlParserADD)|(1<<BeetlParserMIN))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(BeetlParserNOT-40))|(1<<(BeetlParserAT-40))|(1<<(BeetlParserNULL-40))|(1<<(BeetlParserTRUE-40))|(1<<(BeetlParserFALSE-40))|(1<<(BeetlParserEND-40))|(1<<(BeetlParserLEFT_TOKEN-40))|(1<<(BeetlParserLEFT_TOKEN2-40))|(1<<(BeetlParserLEFT_TEXT_TOKEN-40))|(1<<(BeetlParserDecimalLiteral-40))|(1<<(BeetlParserFloatingPointLiteral-40))|(1<<(BeetlParserStringLiteral-40))|(1<<(BeetlParserIdentifier-40))|(1<<(BeetlParserCOMMENT_TAG-40)))) != 0) {
		{
			p.SetState(280)
			p.Statement()
		}

		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVarDeclareListContext is an interface to support dynamic dispatch.
type IVarDeclareListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarDeclareListContext differentiates from other interfaces.
	IsVarDeclareListContext()
}

type VarDeclareListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclareListContext() *VarDeclareListContext {
	var p = new(VarDeclareListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_varDeclareList
	return p
}

func (*VarDeclareListContext) IsVarDeclareListContext() {}

func NewVarDeclareListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclareListContext {
	var p = new(VarDeclareListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_varDeclareList

	return p
}

func (s *VarDeclareListContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclareListContext) AllAssignMent() []IAssignMentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAssignMentContext)(nil)).Elem())
	var tst = make([]IAssignMentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAssignMentContext)
		}
	}

	return tst
}

func (s *VarDeclareListContext) AssignMent(i int) IAssignMentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignMentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAssignMentContext)
}

func (s *VarDeclareListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(BeetlParserCOMMA)
}

func (s *VarDeclareListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(BeetlParserCOMMA, i)
}

func (s *VarDeclareListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclareListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclareListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterVarDeclareList(s)
	}
}

func (s *VarDeclareListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitVarDeclareList(s)
	}
}

func (p *BeetlParser) VarDeclareList() (localctx IVarDeclareListContext) {
	localctx = NewVarDeclareListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BeetlParserRULE_varDeclareList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.AssignMent()
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BeetlParserCOMMA {
		{
			p.SetState(287)
			p.Match(BeetlParserCOMMA)
		}
		{
			p.SetState(288)
			p.AssignMent()
		}

		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAssignMentContext is an interface to support dynamic dispatch.
type IAssignMentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignMentContext differentiates from other interfaces.
	IsAssignMentContext()
}

type AssignMentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignMentContext() *AssignMentContext {
	var p = new(AssignMentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_assignMent
	return p
}

func (*AssignMentContext) IsAssignMentContext() {}

func NewAssignMentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignMentContext {
	var p = new(AssignMentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_assignMent

	return p
}

func (s *AssignMentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignMentContext) CopyFrom(ctx *AssignMentContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AssignMentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignMentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssignGeneralInStContext struct {
	*AssignMentContext
}

func NewAssignGeneralInStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignGeneralInStContext {
	var p = new(AssignGeneralInStContext)

	p.AssignMentContext = NewEmptyAssignMentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AssignMentContext))

	return p
}

func (s *AssignGeneralInStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignGeneralInStContext) GeneralAssignExp() IGeneralAssignExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGeneralAssignExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGeneralAssignExpContext)
}

func (s *AssignGeneralInStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterAssignGeneralInSt(s)
	}
}

func (s *AssignGeneralInStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitAssignGeneralInSt(s)
	}
}

type AssignTemplateVarContext struct {
	*AssignMentContext
}

func NewAssignTemplateVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignTemplateVarContext {
	var p = new(AssignTemplateVarContext)

	p.AssignMentContext = NewEmptyAssignMentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AssignMentContext))

	return p
}

func (s *AssignTemplateVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignTemplateVarContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BeetlParserIdentifier, 0)
}

func (s *AssignTemplateVarContext) ASSIN() antlr.TerminalNode {
	return s.GetToken(BeetlParserASSIN, 0)
}

func (s *AssignTemplateVarContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *AssignTemplateVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterAssignTemplateVar(s)
	}
}

func (s *AssignTemplateVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitAssignTemplateVar(s)
	}
}

type AssignIdContext struct {
	*AssignMentContext
}

func NewAssignIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignIdContext {
	var p = new(AssignIdContext)

	p.AssignMentContext = NewEmptyAssignMentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AssignMentContext))

	return p
}

func (s *AssignIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignIdContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BeetlParserIdentifier, 0)
}

func (s *AssignIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterAssignId(s)
	}
}

func (s *AssignIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitAssignId(s)
	}
}

func (p *BeetlParser) AssignMent() (localctx IAssignMentContext) {
	localctx = NewAssignMentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BeetlParserRULE_assignMent)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignIdContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(294)
			p.Match(BeetlParserIdentifier)
		}

	case 2:
		localctx = NewAssignGeneralInStContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(295)
			p.GeneralAssignExp()
		}

	case 3:
		localctx = NewAssignTemplateVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(296)
			p.Match(BeetlParserIdentifier)
		}
		{
			p.SetState(297)
			p.Match(BeetlParserASSIN)
		}
		{
			p.SetState(298)
			p.Block()
		}

	}

	return localctx
}

// ISwitchBlockContext is an interface to support dynamic dispatch.
type ISwitchBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitchBlockContext differentiates from other interfaces.
	IsSwitchBlockContext()
}

type SwitchBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchBlockContext() *SwitchBlockContext {
	var p = new(SwitchBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_switchBlock
	return p
}

func (*SwitchBlockContext) IsSwitchBlockContext() {}

func NewSwitchBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchBlockContext {
	var p = new(SwitchBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_switchBlock

	return p
}

func (s *SwitchBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchBlockContext) LEFT_BRACE() antlr.TerminalNode {
	return s.GetToken(BeetlParserLEFT_BRACE, 0)
}

func (s *SwitchBlockContext) RIGHT_BRACE() antlr.TerminalNode {
	return s.GetToken(BeetlParserRIGHT_BRACE, 0)
}

func (s *SwitchBlockContext) AllSwitchBlockStatementGroup() []ISwitchBlockStatementGroupContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISwitchBlockStatementGroupContext)(nil)).Elem())
	var tst = make([]ISwitchBlockStatementGroupContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISwitchBlockStatementGroupContext)
		}
	}

	return tst
}

func (s *SwitchBlockContext) SwitchBlockStatementGroup(i int) ISwitchBlockStatementGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitchBlockStatementGroupContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISwitchBlockStatementGroupContext)
}

func (s *SwitchBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterSwitchBlock(s)
	}
}

func (s *SwitchBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitSwitchBlock(s)
	}
}

func (p *BeetlParser) SwitchBlock() (localctx ISwitchBlockContext) {
	localctx = NewSwitchBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, BeetlParserRULE_switchBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		p.Match(BeetlParserLEFT_BRACE)
	}
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BeetlParserCase || _la == BeetlParserDefault {
		{
			p.SetState(302)
			p.SwitchBlockStatementGroup()
		}

		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(308)
		p.Match(BeetlParserRIGHT_BRACE)
	}

	return localctx
}

// ISwitchBlockStatementGroupContext is an interface to support dynamic dispatch.
type ISwitchBlockStatementGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitchBlockStatementGroupContext differentiates from other interfaces.
	IsSwitchBlockStatementGroupContext()
}

type SwitchBlockStatementGroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchBlockStatementGroupContext() *SwitchBlockStatementGroupContext {
	var p = new(SwitchBlockStatementGroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_switchBlockStatementGroup
	return p
}

func (*SwitchBlockStatementGroupContext) IsSwitchBlockStatementGroupContext() {}

func NewSwitchBlockStatementGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchBlockStatementGroupContext {
	var p = new(SwitchBlockStatementGroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_switchBlockStatementGroup

	return p
}

func (s *SwitchBlockStatementGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchBlockStatementGroupContext) AllSwitchLabel() []ISwitchLabelContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISwitchLabelContext)(nil)).Elem())
	var tst = make([]ISwitchLabelContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISwitchLabelContext)
		}
	}

	return tst
}

func (s *SwitchBlockStatementGroupContext) SwitchLabel(i int) ISwitchLabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitchLabelContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISwitchLabelContext)
}

func (s *SwitchBlockStatementGroupContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *SwitchBlockStatementGroupContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *SwitchBlockStatementGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchBlockStatementGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchBlockStatementGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterSwitchBlockStatementGroup(s)
	}
}

func (s *SwitchBlockStatementGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitSwitchBlockStatementGroup(s)
	}
}

func (p *BeetlParser) SwitchBlockStatementGroup() (localctx ISwitchBlockStatementGroupContext) {
	localctx = NewSwitchBlockStatementGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, BeetlParserRULE_switchBlockStatementGroup)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(310)
				p.SwitchLabel()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeetlParserIf)|(1<<BeetlParserFor)|(1<<BeetlParserWhile)|(1<<BeetlParserSwitch)|(1<<BeetlParserSelect)|(1<<BeetlParserReturn)|(1<<BeetlParserBreak)|(1<<BeetlParserContinue)|(1<<BeetlParserVar)|(1<<BeetlParserDirective)|(1<<BeetlParserTry)|(1<<BeetlParserAjax)|(1<<BeetlParserFragment)|(1<<BeetlParserLEFT_BRACE)|(1<<BeetlParserLEFT_PAR)|(1<<BeetlParserLEFT_SQBR)|(1<<BeetlParserINCREASE)|(1<<BeetlParserDECREASE)|(1<<BeetlParserADD)|(1<<BeetlParserMIN))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(BeetlParserNOT-40))|(1<<(BeetlParserAT-40))|(1<<(BeetlParserNULL-40))|(1<<(BeetlParserTRUE-40))|(1<<(BeetlParserFALSE-40))|(1<<(BeetlParserEND-40))|(1<<(BeetlParserLEFT_TOKEN-40))|(1<<(BeetlParserLEFT_TOKEN2-40))|(1<<(BeetlParserLEFT_TEXT_TOKEN-40))|(1<<(BeetlParserDecimalLiteral-40))|(1<<(BeetlParserFloatingPointLiteral-40))|(1<<(BeetlParserStringLiteral-40))|(1<<(BeetlParserIdentifier-40))|(1<<(BeetlParserCOMMENT_TAG-40)))) != 0) {
		{
			p.SetState(315)
			p.Statement()
		}

		p.SetState(320)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISwitchLabelContext is an interface to support dynamic dispatch.
type ISwitchLabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitchLabelContext differentiates from other interfaces.
	IsSwitchLabelContext()
}

type SwitchLabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchLabelContext() *SwitchLabelContext {
	var p = new(SwitchLabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_switchLabel
	return p
}

func (*SwitchLabelContext) IsSwitchLabelContext() {}

func NewSwitchLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchLabelContext {
	var p = new(SwitchLabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_switchLabel

	return p
}

func (s *SwitchLabelContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchLabelContext) Case() antlr.TerminalNode {
	return s.GetToken(BeetlParserCase, 0)
}

func (s *SwitchLabelContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SwitchLabelContext) COLON() antlr.TerminalNode {
	return s.GetToken(BeetlParserCOLON, 0)
}

func (s *SwitchLabelContext) Default() antlr.TerminalNode {
	return s.GetToken(BeetlParserDefault, 0)
}

func (s *SwitchLabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchLabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchLabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterSwitchLabel(s)
	}
}

func (s *SwitchLabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitSwitchLabel(s)
	}
}

func (p *BeetlParser) SwitchLabel() (localctx ISwitchLabelContext) {
	localctx = NewSwitchLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BeetlParserRULE_switchLabel)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(327)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeetlParserCase:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(321)
			p.Match(BeetlParserCase)
		}
		{
			p.SetState(322)
			p.expression(0)
		}
		{
			p.SetState(323)
			p.Match(BeetlParserCOLON)
		}

	case BeetlParserDefault:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(325)
			p.Match(BeetlParserDefault)
		}
		{
			p.SetState(326)
			p.Match(BeetlParserCOLON)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IForControlContext is an interface to support dynamic dispatch.
type IForControlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForControlContext differentiates from other interfaces.
	IsForControlContext()
}

type ForControlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForControlContext() *ForControlContext {
	var p = new(ForControlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_forControl
	return p
}

func (*ForControlContext) IsForControlContext() {}

func NewForControlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForControlContext {
	var p = new(ForControlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_forControl

	return p
}

func (s *ForControlContext) GetParser() antlr.Parser { return s.parser }

func (s *ForControlContext) ForInControl() IForInControlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForInControlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForInControlContext)
}

func (s *ForControlContext) GeneralForControl() IGeneralForControlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGeneralForControlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGeneralForControlContext)
}

func (s *ForControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForControlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForControlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterForControl(s)
	}
}

func (s *ForControlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitForControl(s)
	}
}

func (p *BeetlParser) ForControl() (localctx IForControlContext) {
	localctx = NewForControlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BeetlParserRULE_forControl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(329)
			p.ForInControl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(330)
			p.GeneralForControl()
		}

	}

	return localctx
}

// IForInControlContext is an interface to support dynamic dispatch.
type IForInControlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForInControlContext differentiates from other interfaces.
	IsForInControlContext()
}

type ForInControlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForInControlContext() *ForInControlContext {
	var p = new(ForInControlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_forInControl
	return p
}

func (*ForInControlContext) IsForInControlContext() {}

func NewForInControlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForInControlContext {
	var p = new(ForInControlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_forInControl

	return p
}

func (s *ForInControlContext) GetParser() antlr.Parser { return s.parser }

func (s *ForInControlContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BeetlParserIdentifier, 0)
}

func (s *ForInControlContext) FOR_IN() antlr.TerminalNode {
	return s.GetToken(BeetlParserFOR_IN, 0)
}

func (s *ForInControlContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForInControlContext) Var() antlr.TerminalNode {
	return s.GetToken(BeetlParserVar, 0)
}

func (s *ForInControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInControlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForInControlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterForInControl(s)
	}
}

func (s *ForInControlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitForInControl(s)
	}
}

func (p *BeetlParser) ForInControl() (localctx IForInControlContext) {
	localctx = NewForInControlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, BeetlParserRULE_forInControl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BeetlParserVar {
		{
			p.SetState(333)
			p.Match(BeetlParserVar)
		}

	}
	{
		p.SetState(336)
		p.Match(BeetlParserIdentifier)
	}
	{
		p.SetState(337)
		p.Match(BeetlParserFOR_IN)
	}
	{
		p.SetState(338)
		p.expression(0)
	}

	return localctx
}

// IGeneralForControlContext is an interface to support dynamic dispatch.
type IGeneralForControlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGeneralForControlContext differentiates from other interfaces.
	IsGeneralForControlContext()
}

type GeneralForControlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGeneralForControlContext() *GeneralForControlContext {
	var p = new(GeneralForControlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_generalForControl
	return p
}

func (*GeneralForControlContext) IsGeneralForControlContext() {}

func NewGeneralForControlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeneralForControlContext {
	var p = new(GeneralForControlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_generalForControl

	return p
}

func (s *GeneralForControlContext) GetParser() antlr.Parser { return s.parser }

func (s *GeneralForControlContext) AllEND() []antlr.TerminalNode {
	return s.GetTokens(BeetlParserEND)
}

func (s *GeneralForControlContext) END(i int) antlr.TerminalNode {
	return s.GetToken(BeetlParserEND, i)
}

func (s *GeneralForControlContext) ForInit() IForInitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForInitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForInitContext)
}

func (s *GeneralForControlContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GeneralForControlContext) ForUpdate() IForUpdateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForUpdateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForUpdateContext)
}

func (s *GeneralForControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeneralForControlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeneralForControlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterGeneralForControl(s)
	}
}

func (s *GeneralForControlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitGeneralForControl(s)
	}
}

func (p *BeetlParser) GeneralForControl() (localctx IGeneralForControlContext) {
	localctx = NewGeneralForControlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, BeetlParserRULE_generalForControl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeetlParserVar)|(1<<BeetlParserLEFT_BRACE)|(1<<BeetlParserLEFT_PAR)|(1<<BeetlParserLEFT_SQBR)|(1<<BeetlParserINCREASE)|(1<<BeetlParserDECREASE)|(1<<BeetlParserADD)|(1<<BeetlParserMIN))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(BeetlParserNOT-40))|(1<<(BeetlParserAT-40))|(1<<(BeetlParserNULL-40))|(1<<(BeetlParserTRUE-40))|(1<<(BeetlParserFALSE-40))|(1<<(BeetlParserDecimalLiteral-40))|(1<<(BeetlParserFloatingPointLiteral-40))|(1<<(BeetlParserStringLiteral-40))|(1<<(BeetlParserIdentifier-40)))) != 0) {
		{
			p.SetState(340)
			p.ForInit()
		}

	}
	{
		p.SetState(343)
		p.Match(BeetlParserEND)
	}
	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeetlParserLEFT_BRACE)|(1<<BeetlParserLEFT_PAR)|(1<<BeetlParserLEFT_SQBR)|(1<<BeetlParserINCREASE)|(1<<BeetlParserDECREASE)|(1<<BeetlParserADD)|(1<<BeetlParserMIN))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(BeetlParserNOT-40))|(1<<(BeetlParserAT-40))|(1<<(BeetlParserNULL-40))|(1<<(BeetlParserTRUE-40))|(1<<(BeetlParserFALSE-40))|(1<<(BeetlParserDecimalLiteral-40))|(1<<(BeetlParserFloatingPointLiteral-40))|(1<<(BeetlParserStringLiteral-40))|(1<<(BeetlParserIdentifier-40)))) != 0) {
		{
			p.SetState(344)
			p.expression(0)
		}

	}
	{
		p.SetState(347)
		p.Match(BeetlParserEND)
	}
	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeetlParserLEFT_BRACE)|(1<<BeetlParserLEFT_PAR)|(1<<BeetlParserLEFT_SQBR)|(1<<BeetlParserINCREASE)|(1<<BeetlParserDECREASE)|(1<<BeetlParserADD)|(1<<BeetlParserMIN))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(BeetlParserNOT-40))|(1<<(BeetlParserAT-40))|(1<<(BeetlParserNULL-40))|(1<<(BeetlParserTRUE-40))|(1<<(BeetlParserFALSE-40))|(1<<(BeetlParserDecimalLiteral-40))|(1<<(BeetlParserFloatingPointLiteral-40))|(1<<(BeetlParserStringLiteral-40))|(1<<(BeetlParserIdentifier-40)))) != 0) {
		{
			p.SetState(348)
			p.ForUpdate()
		}

	}

	return localctx
}

// IForInitContext is an interface to support dynamic dispatch.
type IForInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForInitContext differentiates from other interfaces.
	IsForInitContext()
}

type ForInitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForInitContext() *ForInitContext {
	var p = new(ForInitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_forInit
	return p
}

func (*ForInitContext) IsForInitContext() {}

func NewForInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForInitContext {
	var p = new(ForInitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_forInit

	return p
}

func (s *ForInitContext) GetParser() antlr.Parser { return s.parser }

func (s *ForInitContext) Var() antlr.TerminalNode {
	return s.GetToken(BeetlParserVar, 0)
}

func (s *ForInitContext) VarDeclareList() IVarDeclareListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDeclareListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarDeclareListContext)
}

func (s *ForInitContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ForInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterForInit(s)
	}
}

func (s *ForInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitForInit(s)
	}
}

func (p *BeetlParser) ForInit() (localctx IForInitContext) {
	localctx = NewForInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, BeetlParserRULE_forInit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(354)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeetlParserVar:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(351)
			p.Match(BeetlParserVar)
		}
		{
			p.SetState(352)
			p.VarDeclareList()
		}

	case BeetlParserLEFT_BRACE, BeetlParserLEFT_PAR, BeetlParserLEFT_SQBR, BeetlParserINCREASE, BeetlParserDECREASE, BeetlParserADD, BeetlParserMIN, BeetlParserNOT, BeetlParserAT, BeetlParserNULL, BeetlParserTRUE, BeetlParserFALSE, BeetlParserDecimalLiteral, BeetlParserFloatingPointLiteral, BeetlParserStringLiteral, BeetlParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(353)
			p.ExpressionList()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IForUpdateContext is an interface to support dynamic dispatch.
type IForUpdateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForUpdateContext differentiates from other interfaces.
	IsForUpdateContext()
}

type ForUpdateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForUpdateContext() *ForUpdateContext {
	var p = new(ForUpdateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_forUpdate
	return p
}

func (*ForUpdateContext) IsForUpdateContext() {}

func NewForUpdateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForUpdateContext {
	var p = new(ForUpdateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_forUpdate

	return p
}

func (s *ForUpdateContext) GetParser() antlr.Parser { return s.parser }

func (s *ForUpdateContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ForUpdateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForUpdateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForUpdateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterForUpdate(s)
	}
}

func (s *ForUpdateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitForUpdate(s)
	}
}

func (p *BeetlParser) ForUpdate() (localctx IForUpdateContext) {
	localctx = NewForUpdateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, BeetlParserRULE_forUpdate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(356)
		p.ExpressionList()
	}

	return localctx
}

// IParExpressionContext is an interface to support dynamic dispatch.
type IParExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParExpressionContext differentiates from other interfaces.
	IsParExpressionContext()
}

type ParExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParExpressionContext() *ParExpressionContext {
	var p = new(ParExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_parExpression
	return p
}

func (*ParExpressionContext) IsParExpressionContext() {}

func NewParExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParExpressionContext {
	var p = new(ParExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_parExpression

	return p
}

func (s *ParExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ParExpressionContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(BeetlParserLEFT_PAR, 0)
}

func (s *ParExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParExpressionContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(BeetlParserRIGHT_PAR, 0)
}

func (s *ParExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterParExpression(s)
	}
}

func (s *ParExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitParExpression(s)
	}
}

func (p *BeetlParser) ParExpression() (localctx IParExpressionContext) {
	localctx = NewParExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, BeetlParserRULE_parExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(358)
		p.Match(BeetlParserLEFT_PAR)
	}
	{
		p.SetState(359)
		p.expression(0)
	}
	{
		p.SetState(360)
		p.Match(BeetlParserRIGHT_PAR)
	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(BeetlParserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(BeetlParserCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (p *BeetlParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, BeetlParserRULE_expressionList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(362)
		p.expression(0)
	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BeetlParserCOMMA {
		{
			p.SetState(363)
			p.Match(BeetlParserCOMMA)
		}
		{
			p.SetState(364)
			p.expression(0)
		}

		p.SetState(369)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementExpressionContext is an interface to support dynamic dispatch.
type IStatementExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementExpressionContext differentiates from other interfaces.
	IsStatementExpressionContext()
}

type StatementExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementExpressionContext() *StatementExpressionContext {
	var p = new(StatementExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_statementExpression
	return p
}

func (*StatementExpressionContext) IsStatementExpressionContext() {}

func NewStatementExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementExpressionContext {
	var p = new(StatementExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_statementExpression

	return p
}

func (s *StatementExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterStatementExpression(s)
	}
}

func (s *StatementExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitStatementExpression(s)
	}
}

func (p *BeetlParser) StatementExpression() (localctx IStatementExpressionContext) {
	localctx = NewStatementExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, BeetlParserRULE_statementExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		p.expression(0)
	}

	return localctx
}

// ITextStatmentContext is an interface to support dynamic dispatch.
type ITextStatmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTextStatmentContext differentiates from other interfaces.
	IsTextStatmentContext()
}

type TextStatmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextStatmentContext() *TextStatmentContext {
	var p = new(TextStatmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_textStatment
	return p
}

func (*TextStatmentContext) IsTextStatmentContext() {}

func NewTextStatmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextStatmentContext {
	var p = new(TextStatmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_textStatment

	return p
}

func (s *TextStatmentContext) GetParser() antlr.Parser { return s.parser }

func (s *TextStatmentContext) LEFT_TOKEN() antlr.TerminalNode {
	return s.GetToken(BeetlParserLEFT_TOKEN, 0)
}

func (s *TextStatmentContext) TextVar() ITextVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITextVarContext)
}

func (s *TextStatmentContext) RIGHT_TOKEN() antlr.TerminalNode {
	return s.GetToken(BeetlParserRIGHT_TOKEN, 0)
}

func (s *TextStatmentContext) LEFT_TOKEN2() antlr.TerminalNode {
	return s.GetToken(BeetlParserLEFT_TOKEN2, 0)
}

func (s *TextStatmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextStatmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextStatmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterTextStatment(s)
	}
}

func (s *TextStatmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitTextStatment(s)
	}
}

func (p *BeetlParser) TextStatment() (localctx ITextStatmentContext) {
	localctx = NewTextStatmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, BeetlParserRULE_textStatment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(380)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeetlParserLEFT_TOKEN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(372)
			p.Match(BeetlParserLEFT_TOKEN)
		}
		{
			p.SetState(373)
			p.TextVar()
		}
		{
			p.SetState(374)
			p.Match(BeetlParserRIGHT_TOKEN)
		}

	case BeetlParserLEFT_TOKEN2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(376)
			p.Match(BeetlParserLEFT_TOKEN2)
		}
		{
			p.SetState(377)
			p.TextVar()
		}
		{
			p.SetState(378)
			p.Match(BeetlParserRIGHT_TOKEN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITextVarContext is an interface to support dynamic dispatch.
type ITextVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetB returns the b rule contexts.
	GetB() IExpressionContext

	// SetB sets the b rule contexts.
	SetB(IExpressionContext)

	// IsTextVarContext differentiates from other interfaces.
	IsTextVarContext()
}

type TextVarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	b      IExpressionContext
}

func NewEmptyTextVarContext() *TextVarContext {
	var p = new(TextVarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_textVar
	return p
}

func (*TextVarContext) IsTextVarContext() {}

func NewTextVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextVarContext {
	var p = new(TextVarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_textVar

	return p
}

func (s *TextVarContext) GetParser() antlr.Parser { return s.parser }

func (s *TextVarContext) GetB() IExpressionContext { return s.b }

func (s *TextVarContext) SetB(v IExpressionContext) { s.b = v }

func (s *TextVarContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TextVarContext) COMMA() antlr.TerminalNode {
	return s.GetToken(BeetlParserCOMMA, 0)
}

func (s *TextVarContext) Textformat() ITextformatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextformatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITextformatContext)
}

func (s *TextVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterTextVar(s)
	}
}

func (s *TextVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitTextVar(s)
	}
}

func (p *BeetlParser) TextVar() (localctx ITextVarContext) {
	localctx = NewTextVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, BeetlParserRULE_textVar)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(382)

		var _x = p.expression(0)

		localctx.(*TextVarContext).b = _x
	}
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BeetlParserCOMMA {
		{
			p.SetState(383)
			p.Match(BeetlParserCOMMA)
		}
		{
			p.SetState(384)
			p.Textformat()
		}

	}

	return localctx
}

// ITextformatContext is an interface to support dynamic dispatch.
type ITextformatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFm returns the fm rule contexts.
	GetFm() IFunctionNsContext

	// SetFm sets the fm rule contexts.
	SetFm(IFunctionNsContext)

	// IsTextformatContext differentiates from other interfaces.
	IsTextformatContext()
}

type TextformatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	fm     IFunctionNsContext
}

func NewEmptyTextformatContext() *TextformatContext {
	var p = new(TextformatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_textformat
	return p
}

func (*TextformatContext) IsTextformatContext() {}

func NewTextformatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextformatContext {
	var p = new(TextformatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_textformat

	return p
}

func (s *TextformatContext) GetParser() antlr.Parser { return s.parser }

func (s *TextformatContext) GetFm() IFunctionNsContext { return s.fm }

func (s *TextformatContext) SetFm(v IFunctionNsContext) { s.fm = v }

func (s *TextformatContext) FunctionNs() IFunctionNsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionNsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionNsContext)
}

func (s *TextformatContext) ASSIN() antlr.TerminalNode {
	return s.GetToken(BeetlParserASSIN, 0)
}

func (s *TextformatContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(BeetlParserStringLiteral, 0)
}

func (s *TextformatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextformatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextformatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterTextformat(s)
	}
}

func (s *TextformatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitTextformat(s)
	}
}

func (p *BeetlParser) Textformat() (localctx ITextformatContext) {
	localctx = NewTextformatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, BeetlParserRULE_textformat)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(393)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeetlParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(387)

			var _x = p.FunctionNs()

			localctx.(*TextformatContext).fm = _x
		}
		p.SetState(390)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BeetlParserASSIN {
			{
				p.SetState(388)
				p.Match(BeetlParserASSIN)
			}
			{
				p.SetState(389)
				p.Match(BeetlParserStringLiteral)
			}

		}

	case BeetlParserStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(392)
			p.Match(BeetlParserStringLiteral)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConstantsTextStatmentContext is an interface to support dynamic dispatch.
type IConstantsTextStatmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantsTextStatmentContext differentiates from other interfaces.
	IsConstantsTextStatmentContext()
}

type ConstantsTextStatmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantsTextStatmentContext() *ConstantsTextStatmentContext {
	var p = new(ConstantsTextStatmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_constantsTextStatment
	return p
}

func (*ConstantsTextStatmentContext) IsConstantsTextStatmentContext() {}

func NewConstantsTextStatmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantsTextStatmentContext {
	var p = new(ConstantsTextStatmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_constantsTextStatment

	return p
}

func (s *ConstantsTextStatmentContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantsTextStatmentContext) LEFT_TEXT_TOKEN() antlr.TerminalNode {
	return s.GetToken(BeetlParserLEFT_TEXT_TOKEN, 0)
}

func (s *ConstantsTextStatmentContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(BeetlParserDecimalLiteral, 0)
}

func (s *ConstantsTextStatmentContext) RIGHT_TOKEN() antlr.TerminalNode {
	return s.GetToken(BeetlParserRIGHT_TOKEN, 0)
}

func (s *ConstantsTextStatmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantsTextStatmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantsTextStatmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterConstantsTextStatment(s)
	}
}

func (s *ConstantsTextStatmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitConstantsTextStatment(s)
	}
}

func (p *BeetlParser) ConstantsTextStatment() (localctx IConstantsTextStatmentContext) {
	localctx = NewConstantsTextStatmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, BeetlParserRULE_constantsTextStatment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(395)
		p.Match(BeetlParserLEFT_TEXT_TOKEN)
	}
	{
		p.SetState(396)
		p.Match(BeetlParserDecimalLiteral)
	}
	{
		p.SetState(397)
		p.Match(BeetlParserRIGHT_TOKEN)
	}

	return localctx
}

// IConstantExpressionContext is an interface to support dynamic dispatch.
type IConstantExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantExpressionContext differentiates from other interfaces.
	IsConstantExpressionContext()
}

type ConstantExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantExpressionContext() *ConstantExpressionContext {
	var p = new(ConstantExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_constantExpression
	return p
}

func (*ConstantExpressionContext) IsConstantExpressionContext() {}

func NewConstantExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantExpressionContext {
	var p = new(ConstantExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_constantExpression

	return p
}

func (s *ConstantExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConstantExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterConstantExpression(s)
	}
}

func (s *ConstantExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitConstantExpression(s)
	}
}

func (p *BeetlParser) ConstantExpression() (localctx IConstantExpressionContext) {
	localctx = NewConstantExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, BeetlParserRULE_constantExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(399)
		p.expression(0)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IncDecOneContext struct {
	*ExpressionContext
}

func NewIncDecOneContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IncDecOneContext {
	var p = new(IncDecOneContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IncDecOneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncDecOneContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BeetlParserIdentifier, 0)
}

func (s *IncDecOneContext) INCREASE() antlr.TerminalNode {
	return s.GetToken(BeetlParserINCREASE, 0)
}

func (s *IncDecOneContext) DECREASE() antlr.TerminalNode {
	return s.GetToken(BeetlParserDECREASE, 0)
}

func (s *IncDecOneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterIncDecOne(s)
	}
}

func (s *IncDecOneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitIncDecOne(s)
	}
}

type AddminExpContext struct {
	*ExpressionContext
}

func NewAddminExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddminExpContext {
	var p = new(AddminExpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AddminExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddminExpContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AddminExpContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddminExpContext) ADD() antlr.TerminalNode {
	return s.GetToken(BeetlParserADD, 0)
}

func (s *AddminExpContext) MIN() antlr.TerminalNode {
	return s.GetToken(BeetlParserMIN, 0)
}

func (s *AddminExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterAddminExp(s)
	}
}

func (s *AddminExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitAddminExp(s)
	}
}

type NativeCallExpContext struct {
	*ExpressionContext
}

func NewNativeCallExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NativeCallExpContext {
	var p = new(NativeCallExpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NativeCallExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NativeCallExpContext) AT() antlr.TerminalNode {
	return s.GetToken(BeetlParserAT, 0)
}

func (s *NativeCallExpContext) NativeCall() INativeCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INativeCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INativeCallContext)
}

func (s *NativeCallExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterNativeCallExp(s)
	}
}

func (s *NativeCallExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitNativeCallExp(s)
	}
}

type AndExpContext struct {
	*ExpressionContext
}

func NewAndExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExpContext {
	var p = new(AndExpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AndExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AndExpContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AndExpContext) AND() antlr.TerminalNode {
	return s.GetToken(BeetlParserAND, 0)
}

func (s *AndExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterAndExp(s)
	}
}

func (s *AndExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitAndExp(s)
	}
}

type FunctionCallExpContext struct {
	*ExpressionContext
}

func NewFunctionCallExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallExpContext {
	var p = new(FunctionCallExpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *FunctionCallExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallExpContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *FunctionCallExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterFunctionCallExp(s)
	}
}

func (s *FunctionCallExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitFunctionCallExp(s)
	}
}

type AssignGeneralInExpContext struct {
	*ExpressionContext
}

func NewAssignGeneralInExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignGeneralInExpContext {
	var p = new(AssignGeneralInExpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AssignGeneralInExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignGeneralInExpContext) GeneralAssignExp() IGeneralAssignExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGeneralAssignExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGeneralAssignExpContext)
}

func (s *AssignGeneralInExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterAssignGeneralInExp(s)
	}
}

func (s *AssignGeneralInExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitAssignGeneralInExp(s)
	}
}

type OrExpContext struct {
	*ExpressionContext
}

func NewOrExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExpContext {
	var p = new(OrExpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *OrExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExpContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *OrExpContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OrExpContext) OR() antlr.TerminalNode {
	return s.GetToken(BeetlParserOR, 0)
}

func (s *OrExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterOrExp(s)
	}
}

func (s *OrExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitOrExp(s)
	}
}

type NotExpContext struct {
	*ExpressionContext
}

func NewNotExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExpContext {
	var p = new(NotExpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NotExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExpContext) NOT() antlr.TerminalNode {
	return s.GetToken(BeetlParserNOT, 0)
}

func (s *NotExpContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NotExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterNotExp(s)
	}
}

func (s *NotExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitNotExp(s)
	}
}

type MuldivmodExpContext struct {
	*ExpressionContext
}

func NewMuldivmodExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MuldivmodExpContext {
	var p = new(MuldivmodExpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MuldivmodExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MuldivmodExpContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MuldivmodExpContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MuldivmodExpContext) MUlTIP() antlr.TerminalNode {
	return s.GetToken(BeetlParserMUlTIP, 0)
}

func (s *MuldivmodExpContext) DIV() antlr.TerminalNode {
	return s.GetToken(BeetlParserDIV, 0)
}

func (s *MuldivmodExpContext) MOD() antlr.TerminalNode {
	return s.GetToken(BeetlParserMOD, 0)
}

func (s *MuldivmodExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterMuldivmodExp(s)
	}
}

func (s *MuldivmodExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitMuldivmodExp(s)
	}
}

type CompareExpContext struct {
	*ExpressionContext
}

func NewCompareExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareExpContext {
	var p = new(CompareExpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CompareExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareExpContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *CompareExpContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CompareExpContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(BeetlParserEQUAL, 0)
}

func (s *CompareExpContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(BeetlParserNOT_EQUAL, 0)
}

func (s *CompareExpContext) LESS_EQUAL() antlr.TerminalNode {
	return s.GetToken(BeetlParserLESS_EQUAL, 0)
}

func (s *CompareExpContext) LARGE_EQUAL() antlr.TerminalNode {
	return s.GetToken(BeetlParserLARGE_EQUAL, 0)
}

func (s *CompareExpContext) LARGE() antlr.TerminalNode {
	return s.GetToken(BeetlParserLARGE, 0)
}

func (s *CompareExpContext) LESS() antlr.TerminalNode {
	return s.GetToken(BeetlParserLESS, 0)
}

func (s *CompareExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterCompareExp(s)
	}
}

func (s *CompareExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitCompareExp(s)
	}
}

type LiteralExpContext struct {
	*ExpressionContext
}

func NewLiteralExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpContext {
	var p = new(LiteralExpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LiteralExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterLiteralExp(s)
	}
}

func (s *LiteralExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitLiteralExp(s)
	}
}

type JsonExpContext struct {
	*ExpressionContext
}

func NewJsonExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonExpContext {
	var p = new(JsonExpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *JsonExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonExpContext) Json() IJsonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonContext)
}

func (s *JsonExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterJsonExp(s)
	}
}

func (s *JsonExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitJsonExp(s)
	}
}

type ParExpContext struct {
	*ExpressionContext
}

func NewParExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParExpContext {
	var p = new(ParExpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ParExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParExpContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(BeetlParserLEFT_PAR, 0)
}

func (s *ParExpContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParExpContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(BeetlParserRIGHT_PAR, 0)
}

func (s *ParExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterParExp(s)
	}
}

func (s *ParExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitParExp(s)
	}
}

type NegExpContext struct {
	*ExpressionContext
}

func NewNegExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegExpContext {
	var p = new(NegExpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NegExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegExpContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NegExpContext) ADD() antlr.TerminalNode {
	return s.GetToken(BeetlParserADD, 0)
}

func (s *NegExpContext) MIN() antlr.TerminalNode {
	return s.GetToken(BeetlParserMIN, 0)
}

func (s *NegExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterNegExp(s)
	}
}

func (s *NegExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitNegExp(s)
	}
}

type OneIncDecContext struct {
	*ExpressionContext
}

func NewOneIncDecContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OneIncDecContext {
	var p = new(OneIncDecContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *OneIncDecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneIncDecContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BeetlParserIdentifier, 0)
}

func (s *OneIncDecContext) INCREASE() antlr.TerminalNode {
	return s.GetToken(BeetlParserINCREASE, 0)
}

func (s *OneIncDecContext) DECREASE() antlr.TerminalNode {
	return s.GetToken(BeetlParserDECREASE, 0)
}

func (s *OneIncDecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterOneIncDec(s)
	}
}

func (s *OneIncDecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitOneIncDec(s)
	}
}

type TernaryExpContext struct {
	*ExpressionContext
}

func NewTernaryExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TernaryExpContext {
	var p = new(TernaryExpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *TernaryExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryExpContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *TernaryExpContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TernaryExpContext) QUESTOIN() antlr.TerminalNode {
	return s.GetToken(BeetlParserQUESTOIN, 0)
}

func (s *TernaryExpContext) COLON() antlr.TerminalNode {
	return s.GetToken(BeetlParserCOLON, 0)
}

func (s *TernaryExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterTernaryExp(s)
	}
}

func (s *TernaryExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitTernaryExp(s)
	}
}

type VarRefExpContext struct {
	*ExpressionContext
}

func NewVarRefExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarRefExpContext {
	var p = new(VarRefExpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *VarRefExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarRefExpContext) VarRef() IVarRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarRefContext)
}

func (s *VarRefExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterVarRefExp(s)
	}
}

func (s *VarRefExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitVarRefExp(s)
	}
}

type AssingSelfExpContext struct {
	*ExpressionContext
}

func NewAssingSelfExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssingSelfExpContext {
	var p = new(AssingSelfExpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AssingSelfExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssingSelfExpContext) GeneralAssingSelfExp() IGeneralAssingSelfExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGeneralAssingSelfExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGeneralAssingSelfExpContext)
}

func (s *AssingSelfExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterAssingSelfExp(s)
	}
}

func (s *AssingSelfExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitAssingSelfExp(s)
	}
}

func (p *BeetlParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *BeetlParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 62
	p.EnterRecursionRule(localctx, 62, BeetlParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLiteralExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(402)
			p.Literal()
		}

	case 2:
		localctx = NewNativeCallExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(403)
			p.Match(BeetlParserAT)
		}
		{
			p.SetState(404)
			p.NativeCall()
		}

	case 3:
		localctx = NewFunctionCallExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(405)
			p.FunctionCall()
		}

	case 4:
		localctx = NewVarRefExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(406)
			p.VarRef()
		}

	case 5:
		localctx = NewJsonExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(407)
			p.Json()
		}

	case 6:
		localctx = NewOneIncDecContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(408)
			p.Match(BeetlParserIdentifier)
		}
		{
			p.SetState(409)
			_la = p.GetTokenStream().LA(1)

			if !(_la == BeetlParserINCREASE || _la == BeetlParserDECREASE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 7:
		localctx = NewNegExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(410)
			_la = p.GetTokenStream().LA(1)

			if !(_la == BeetlParserADD || _la == BeetlParserMIN) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(411)
			p.expression(12)
		}

	case 8:
		localctx = NewIncDecOneContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(412)
			_la = p.GetTokenStream().LA(1)

			if !(_la == BeetlParserINCREASE || _la == BeetlParserDECREASE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(413)
			p.Match(BeetlParserIdentifier)
		}

	case 9:
		localctx = NewNotExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(414)
			p.Match(BeetlParserNOT)
		}
		{
			p.SetState(415)
			p.expression(10)
		}

	case 10:
		localctx = NewParExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(416)
			p.Match(BeetlParserLEFT_PAR)
		}
		{
			p.SetState(417)
			p.expression(0)
		}
		{
			p.SetState(418)
			p.Match(BeetlParserRIGHT_PAR)
		}

	case 11:
		localctx = NewAssignGeneralInExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(420)
			p.GeneralAssignExp()
		}

	case 12:
		localctx = NewAssingSelfExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(421)
			p.GeneralAssingSelfExp()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(452)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(450)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMuldivmodExpContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BeetlParserRULE_expression)
				p.SetState(424)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(425)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-30)&-(0x1f+1)) == 0 && ((1<<uint((_la-30)))&((1<<(BeetlParserMUlTIP-30))|(1<<(BeetlParserDIV-30))|(1<<(BeetlParserMOD-30)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(426)
					p.expression(10)
				}

			case 2:
				localctx = NewAddminExpContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BeetlParserRULE_expression)
				p.SetState(427)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(428)
					_la = p.GetTokenStream().LA(1)

					if !(_la == BeetlParserADD || _la == BeetlParserMIN) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(429)
					p.expression(9)
				}

			case 3:
				localctx = NewCompareExpContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BeetlParserRULE_expression)
				p.SetState(430)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(431)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(BeetlParserEQUAL-33))|(1<<(BeetlParserNOT_EQUAL-33))|(1<<(BeetlParserLARGE_EQUAL-33))|(1<<(BeetlParserLARGE-33))|(1<<(BeetlParserLESS_EQUAL-33))|(1<<(BeetlParserLESS-33)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(432)
					p.expression(8)
				}

			case 4:
				localctx = NewAndExpContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BeetlParserRULE_expression)
				p.SetState(433)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(434)
					p.Match(BeetlParserAND)
				}
				{
					p.SetState(435)
					p.expression(7)
				}

			case 5:
				localctx = NewOrExpContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BeetlParserRULE_expression)
				p.SetState(436)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(437)
					p.Match(BeetlParserOR)
				}
				{
					p.SetState(438)
					p.expression(6)
				}

			case 6:
				localctx = NewTernaryExpContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BeetlParserRULE_expression)
				p.SetState(439)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(440)
					p.Match(BeetlParserQUESTOIN)
				}
				p.SetState(442)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(441)
						p.expression(0)
					}

				}
				p.SetState(445)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(444)
						p.Match(BeetlParserCOLON)
					}

				}
				p.SetState(448)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(447)
						p.expression(0)
					}

				}

			}

		}
		p.SetState(454)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())
	}

	return localctx
}

// IGeneralAssignExpContext is an interface to support dynamic dispatch.
type IGeneralAssignExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGeneralAssignExpContext differentiates from other interfaces.
	IsGeneralAssignExpContext()
}

type GeneralAssignExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGeneralAssignExpContext() *GeneralAssignExpContext {
	var p = new(GeneralAssignExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_generalAssignExp
	return p
}

func (*GeneralAssignExpContext) IsGeneralAssignExpContext() {}

func NewGeneralAssignExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeneralAssignExpContext {
	var p = new(GeneralAssignExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_generalAssignExp

	return p
}

func (s *GeneralAssignExpContext) GetParser() antlr.Parser { return s.parser }

func (s *GeneralAssignExpContext) VarRef() IVarRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarRefContext)
}

func (s *GeneralAssignExpContext) ASSIN() antlr.TerminalNode {
	return s.GetToken(BeetlParserASSIN, 0)
}

func (s *GeneralAssignExpContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GeneralAssignExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeneralAssignExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeneralAssignExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterGeneralAssignExp(s)
	}
}

func (s *GeneralAssignExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitGeneralAssignExp(s)
	}
}

func (p *BeetlParser) GeneralAssignExp() (localctx IGeneralAssignExpContext) {
	localctx = NewGeneralAssignExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, BeetlParserRULE_generalAssignExp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(455)
		p.VarRef()
	}
	{
		p.SetState(456)
		p.Match(BeetlParserASSIN)
	}
	{
		p.SetState(457)
		p.expression(0)
	}

	return localctx
}

// IVarRefContext is an interface to support dynamic dispatch.
type IVarRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarRefContext differentiates from other interfaces.
	IsVarRefContext()
}

type VarRefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarRefContext() *VarRefContext {
	var p = new(VarRefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_varRef
	return p
}

func (*VarRefContext) IsVarRefContext() {}

func NewVarRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarRefContext {
	var p = new(VarRefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_varRef

	return p
}

func (s *VarRefContext) GetParser() antlr.Parser { return s.parser }

func (s *VarRefContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BeetlParserIdentifier, 0)
}

func (s *VarRefContext) AllVarAttribute() []IVarAttributeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVarAttributeContext)(nil)).Elem())
	var tst = make([]IVarAttributeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVarAttributeContext)
		}
	}

	return tst
}

func (s *VarRefContext) VarAttribute(i int) IVarAttributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarAttributeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVarAttributeContext)
}

func (s *VarRefContext) Safe_output() ISafe_outputContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISafe_outputContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISafe_outputContext)
}

func (s *VarRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterVarRef(s)
	}
}

func (s *VarRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitVarRef(s)
	}
}

func (p *BeetlParser) VarRef() (localctx IVarRefContext) {
	localctx = NewVarRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, BeetlParserRULE_varRef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(459)
		p.Match(BeetlParserIdentifier)
	}
	p.SetState(463)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(460)
				p.VarAttribute()
			}

		}
		p.SetState(465)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())
	}
	p.SetState(467)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(466)
			p.Safe_output()
		}

	}

	return localctx
}

// IGeneralAssingSelfExpContext is an interface to support dynamic dispatch.
type IGeneralAssingSelfExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGeneralAssingSelfExpContext differentiates from other interfaces.
	IsGeneralAssingSelfExpContext()
}

type GeneralAssingSelfExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGeneralAssingSelfExpContext() *GeneralAssingSelfExpContext {
	var p = new(GeneralAssingSelfExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_generalAssingSelfExp
	return p
}

func (*GeneralAssingSelfExpContext) IsGeneralAssingSelfExpContext() {}

func NewGeneralAssingSelfExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeneralAssingSelfExpContext {
	var p = new(GeneralAssingSelfExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_generalAssingSelfExp

	return p
}

func (s *GeneralAssingSelfExpContext) GetParser() antlr.Parser { return s.parser }

func (s *GeneralAssingSelfExpContext) VarRef() IVarRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarRefContext)
}

func (s *GeneralAssingSelfExpContext) ASSIN() antlr.TerminalNode {
	return s.GetToken(BeetlParserASSIN, 0)
}

func (s *GeneralAssingSelfExpContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GeneralAssingSelfExpContext) ADD() antlr.TerminalNode {
	return s.GetToken(BeetlParserADD, 0)
}

func (s *GeneralAssingSelfExpContext) MIN() antlr.TerminalNode {
	return s.GetToken(BeetlParserMIN, 0)
}

func (s *GeneralAssingSelfExpContext) MUlTIP() antlr.TerminalNode {
	return s.GetToken(BeetlParserMUlTIP, 0)
}

func (s *GeneralAssingSelfExpContext) DIV() antlr.TerminalNode {
	return s.GetToken(BeetlParserDIV, 0)
}

func (s *GeneralAssingSelfExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeneralAssingSelfExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeneralAssingSelfExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterGeneralAssingSelfExp(s)
	}
}

func (s *GeneralAssingSelfExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitGeneralAssingSelfExp(s)
	}
}

func (p *BeetlParser) GeneralAssingSelfExp() (localctx IGeneralAssingSelfExpContext) {
	localctx = NewGeneralAssingSelfExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, BeetlParserRULE_generalAssingSelfExp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(469)
		p.VarRef()
	}
	{
		p.SetState(470)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeetlParserADD)|(1<<BeetlParserMIN)|(1<<BeetlParserMUlTIP)|(1<<BeetlParserDIV))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(471)
		p.Match(BeetlParserASSIN)
	}
	{
		p.SetState(472)
		p.expression(0)
	}

	return localctx
}

// IVarAttributeContext is an interface to support dynamic dispatch.
type IVarAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarAttributeContext differentiates from other interfaces.
	IsVarAttributeContext()
}

type VarAttributeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarAttributeContext() *VarAttributeContext {
	var p = new(VarAttributeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_varAttribute
	return p
}

func (*VarAttributeContext) IsVarAttributeContext() {}

func NewVarAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarAttributeContext {
	var p = new(VarAttributeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_varAttribute

	return p
}

func (s *VarAttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *VarAttributeContext) CopyFrom(ctx *VarAttributeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *VarAttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarAttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VarAttributeVirtualContext struct {
	*VarAttributeContext
}

func NewVarAttributeVirtualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarAttributeVirtualContext {
	var p = new(VarAttributeVirtualContext)

	p.VarAttributeContext = NewEmptyVarAttributeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*VarAttributeContext))

	return p
}

func (s *VarAttributeVirtualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarAttributeVirtualContext) VIRTUAL() antlr.TerminalNode {
	return s.GetToken(BeetlParserVIRTUAL, 0)
}

func (s *VarAttributeVirtualContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BeetlParserIdentifier, 0)
}

func (s *VarAttributeVirtualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterVarAttributeVirtual(s)
	}
}

func (s *VarAttributeVirtualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitVarAttributeVirtual(s)
	}
}

type VarAttributeGeneralContext struct {
	*VarAttributeContext
}

func NewVarAttributeGeneralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarAttributeGeneralContext {
	var p = new(VarAttributeGeneralContext)

	p.VarAttributeContext = NewEmptyVarAttributeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*VarAttributeContext))

	return p
}

func (s *VarAttributeGeneralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarAttributeGeneralContext) PERIOD() antlr.TerminalNode {
	return s.GetToken(BeetlParserPERIOD, 0)
}

func (s *VarAttributeGeneralContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BeetlParserIdentifier, 0)
}

func (s *VarAttributeGeneralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterVarAttributeGeneral(s)
	}
}

func (s *VarAttributeGeneralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitVarAttributeGeneral(s)
	}
}

type VarAttributeArrayOrMapContext struct {
	*VarAttributeContext
}

func NewVarAttributeArrayOrMapContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarAttributeArrayOrMapContext {
	var p = new(VarAttributeArrayOrMapContext)

	p.VarAttributeContext = NewEmptyVarAttributeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*VarAttributeContext))

	return p
}

func (s *VarAttributeArrayOrMapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarAttributeArrayOrMapContext) LEFT_SQBR() antlr.TerminalNode {
	return s.GetToken(BeetlParserLEFT_SQBR, 0)
}

func (s *VarAttributeArrayOrMapContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VarAttributeArrayOrMapContext) RIGHT_SQBR() antlr.TerminalNode {
	return s.GetToken(BeetlParserRIGHT_SQBR, 0)
}

func (s *VarAttributeArrayOrMapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterVarAttributeArrayOrMap(s)
	}
}

func (s *VarAttributeArrayOrMapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitVarAttributeArrayOrMap(s)
	}
}

func (p *BeetlParser) VarAttribute() (localctx IVarAttributeContext) {
	localctx = NewVarAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, BeetlParserRULE_varAttribute)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(482)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeetlParserPERIOD:
		localctx = NewVarAttributeGeneralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(474)
			p.Match(BeetlParserPERIOD)
		}
		{
			p.SetState(475)
			p.Match(BeetlParserIdentifier)
		}

	case BeetlParserVIRTUAL:
		localctx = NewVarAttributeVirtualContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(476)
			p.Match(BeetlParserVIRTUAL)
		}
		{
			p.SetState(477)
			p.Match(BeetlParserIdentifier)
		}

	case BeetlParserLEFT_SQBR:
		localctx = NewVarAttributeArrayOrMapContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(478)
			p.Match(BeetlParserLEFT_SQBR)
		}
		{
			p.SetState(479)
			p.expression(0)
		}
		{
			p.SetState(480)
			p.Match(BeetlParserRIGHT_SQBR)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISafe_outputContext is an interface to support dynamic dispatch.
type ISafe_outputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSafe_outputContext differentiates from other interfaces.
	IsSafe_outputContext()
}

type Safe_outputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySafe_outputContext() *Safe_outputContext {
	var p = new(Safe_outputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_safe_output
	return p
}

func (*Safe_outputContext) IsSafe_outputContext() {}

func NewSafe_outputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Safe_outputContext {
	var p = new(Safe_outputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_safe_output

	return p
}

func (s *Safe_outputContext) GetParser() antlr.Parser { return s.parser }

func (s *Safe_outputContext) NOT() antlr.TerminalNode {
	return s.GetToken(BeetlParserNOT, 0)
}

func (s *Safe_outputContext) Safe_allow_exp() ISafe_allow_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISafe_allow_expContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISafe_allow_expContext)
}

func (s *Safe_outputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Safe_outputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Safe_outputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterSafe_output(s)
	}
}

func (s *Safe_outputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitSafe_output(s)
	}
}

func (p *BeetlParser) Safe_output() (localctx ISafe_outputContext) {
	localctx = NewSafe_outputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, BeetlParserRULE_safe_output)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(484)
		p.Match(BeetlParserNOT)
	}
	p.SetState(486)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(485)
			p.Safe_allow_exp()
		}

	}

	return localctx
}

// ISafe_allow_expContext is an interface to support dynamic dispatch.
type ISafe_allow_expContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSafe_allow_expContext differentiates from other interfaces.
	IsSafe_allow_expContext()
}

type Safe_allow_expContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySafe_allow_expContext() *Safe_allow_expContext {
	var p = new(Safe_allow_expContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_safe_allow_exp
	return p
}

func (*Safe_allow_expContext) IsSafe_allow_expContext() {}

func NewSafe_allow_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Safe_allow_expContext {
	var p = new(Safe_allow_expContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_safe_allow_exp

	return p
}

func (s *Safe_allow_expContext) GetParser() antlr.Parser { return s.parser }

func (s *Safe_allow_expContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *Safe_allow_expContext) AT() antlr.TerminalNode {
	return s.GetToken(BeetlParserAT, 0)
}

func (s *Safe_allow_expContext) NativeCall() INativeCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INativeCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INativeCallContext)
}

func (s *Safe_allow_expContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *Safe_allow_expContext) Json() IJsonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonContext)
}

func (s *Safe_allow_expContext) VarRef() IVarRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarRefContext)
}

func (s *Safe_allow_expContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(BeetlParserLEFT_PAR, 0)
}

func (s *Safe_allow_expContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Safe_allow_expContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(BeetlParserRIGHT_PAR, 0)
}

func (s *Safe_allow_expContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Safe_allow_expContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Safe_allow_expContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterSafe_allow_exp(s)
	}
}

func (s *Safe_allow_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitSafe_allow_exp(s)
	}
}

func (p *BeetlParser) Safe_allow_exp() (localctx ISafe_allow_expContext) {
	localctx = NewSafe_allow_expContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, BeetlParserRULE_safe_allow_exp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(498)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(488)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(489)
			p.Match(BeetlParserAT)
		}
		{
			p.SetState(490)
			p.NativeCall()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(491)
			p.FunctionCall()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(492)
			p.Json()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(493)
			p.VarRef()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(494)
			p.Match(BeetlParserLEFT_PAR)
		}
		{
			p.SetState(495)
			p.expression(0)
		}
		{
			p.SetState(496)
			p.Match(BeetlParserRIGHT_PAR)
		}

	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) FunctionNs() IFunctionNsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionNsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionNsContext)
}

func (s *FunctionCallContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(BeetlParserLEFT_PAR, 0)
}

func (s *FunctionCallContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(BeetlParserRIGHT_PAR, 0)
}

func (s *FunctionCallContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *FunctionCallContext) AllVarAttribute() []IVarAttributeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVarAttributeContext)(nil)).Elem())
	var tst = make([]IVarAttributeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVarAttributeContext)
		}
	}

	return tst
}

func (s *FunctionCallContext) VarAttribute(i int) IVarAttributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarAttributeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVarAttributeContext)
}

func (s *FunctionCallContext) Safe_output() ISafe_outputContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISafe_outputContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISafe_outputContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *BeetlParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, BeetlParserRULE_functionCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(500)
		p.FunctionNs()
	}
	{
		p.SetState(501)
		p.Match(BeetlParserLEFT_PAR)
	}
	p.SetState(503)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeetlParserLEFT_BRACE)|(1<<BeetlParserLEFT_PAR)|(1<<BeetlParserLEFT_SQBR)|(1<<BeetlParserINCREASE)|(1<<BeetlParserDECREASE)|(1<<BeetlParserADD)|(1<<BeetlParserMIN))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(BeetlParserNOT-40))|(1<<(BeetlParserAT-40))|(1<<(BeetlParserNULL-40))|(1<<(BeetlParserTRUE-40))|(1<<(BeetlParserFALSE-40))|(1<<(BeetlParserDecimalLiteral-40))|(1<<(BeetlParserFloatingPointLiteral-40))|(1<<(BeetlParserStringLiteral-40))|(1<<(BeetlParserIdentifier-40)))) != 0) {
		{
			p.SetState(502)
			p.ExpressionList()
		}

	}
	{
		p.SetState(505)
		p.Match(BeetlParserRIGHT_PAR)
	}
	p.SetState(509)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(506)
				p.VarAttribute()
			}

		}
		p.SetState(511)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext())
	}
	p.SetState(513)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(512)
			p.Safe_output()
		}

	}

	return localctx
}

// IFunctionTagCallContext is an interface to support dynamic dispatch.
type IFunctionTagCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionTagCallContext differentiates from other interfaces.
	IsFunctionTagCallContext()
}

type FunctionTagCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionTagCallContext() *FunctionTagCallContext {
	var p = new(FunctionTagCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_functionTagCall
	return p
}

func (*FunctionTagCallContext) IsFunctionTagCallContext() {}

func NewFunctionTagCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionTagCallContext {
	var p = new(FunctionTagCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_functionTagCall

	return p
}

func (s *FunctionTagCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionTagCallContext) FunctionNs() IFunctionNsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionNsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionNsContext)
}

func (s *FunctionTagCallContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(BeetlParserLEFT_PAR, 0)
}

func (s *FunctionTagCallContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(BeetlParserRIGHT_PAR, 0)
}

func (s *FunctionTagCallContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionTagCallContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *FunctionTagCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionTagCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionTagCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterFunctionTagCall(s)
	}
}

func (s *FunctionTagCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitFunctionTagCall(s)
	}
}

func (p *BeetlParser) FunctionTagCall() (localctx IFunctionTagCallContext) {
	localctx = NewFunctionTagCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, BeetlParserRULE_functionTagCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(515)
		p.FunctionNs()
	}
	{
		p.SetState(516)
		p.Match(BeetlParserLEFT_PAR)
	}
	p.SetState(518)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeetlParserLEFT_BRACE)|(1<<BeetlParserLEFT_PAR)|(1<<BeetlParserLEFT_SQBR)|(1<<BeetlParserINCREASE)|(1<<BeetlParserDECREASE)|(1<<BeetlParserADD)|(1<<BeetlParserMIN))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(BeetlParserNOT-40))|(1<<(BeetlParserAT-40))|(1<<(BeetlParserNULL-40))|(1<<(BeetlParserTRUE-40))|(1<<(BeetlParserFALSE-40))|(1<<(BeetlParserDecimalLiteral-40))|(1<<(BeetlParserFloatingPointLiteral-40))|(1<<(BeetlParserStringLiteral-40))|(1<<(BeetlParserIdentifier-40)))) != 0) {
		{
			p.SetState(517)
			p.ExpressionList()
		}

	}
	{
		p.SetState(520)
		p.Match(BeetlParserRIGHT_PAR)
	}
	{
		p.SetState(521)
		p.Block()
	}

	return localctx
}

// IFunctionNsContext is an interface to support dynamic dispatch.
type IFunctionNsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionNsContext differentiates from other interfaces.
	IsFunctionNsContext()
}

type FunctionNsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionNsContext() *FunctionNsContext {
	var p = new(FunctionNsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_functionNs
	return p
}

func (*FunctionNsContext) IsFunctionNsContext() {}

func NewFunctionNsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionNsContext {
	var p = new(FunctionNsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_functionNs

	return p
}

func (s *FunctionNsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionNsContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(BeetlParserIdentifier)
}

func (s *FunctionNsContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(BeetlParserIdentifier, i)
}

func (s *FunctionNsContext) AllPERIOD() []antlr.TerminalNode {
	return s.GetTokens(BeetlParserPERIOD)
}

func (s *FunctionNsContext) PERIOD(i int) antlr.TerminalNode {
	return s.GetToken(BeetlParserPERIOD, i)
}

func (s *FunctionNsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionNsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionNsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterFunctionNs(s)
	}
}

func (s *FunctionNsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitFunctionNs(s)
	}
}

func (p *BeetlParser) FunctionNs() (localctx IFunctionNsContext) {
	localctx = NewFunctionNsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, BeetlParserRULE_functionNs)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(523)
		p.Match(BeetlParserIdentifier)
	}
	p.SetState(528)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BeetlParserPERIOD {
		{
			p.SetState(524)
			p.Match(BeetlParserPERIOD)
		}
		{
			p.SetState(525)
			p.Match(BeetlParserIdentifier)
		}

		p.SetState(530)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INativeCallContext is an interface to support dynamic dispatch.
type INativeCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNativeCallContext differentiates from other interfaces.
	IsNativeCallContext()
}

type NativeCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNativeCallContext() *NativeCallContext {
	var p = new(NativeCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_nativeCall
	return p
}

func (*NativeCallContext) IsNativeCallContext() {}

func NewNativeCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NativeCallContext {
	var p = new(NativeCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_nativeCall

	return p
}

func (s *NativeCallContext) GetParser() antlr.Parser { return s.parser }

func (s *NativeCallContext) AllNativeVarRefChain() []INativeVarRefChainContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INativeVarRefChainContext)(nil)).Elem())
	var tst = make([]INativeVarRefChainContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INativeVarRefChainContext)
		}
	}

	return tst
}

func (s *NativeCallContext) NativeVarRefChain(i int) INativeVarRefChainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INativeVarRefChainContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INativeVarRefChainContext)
}

func (s *NativeCallContext) AllNativeMethod() []INativeMethodContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INativeMethodContext)(nil)).Elem())
	var tst = make([]INativeMethodContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INativeMethodContext)
		}
	}

	return tst
}

func (s *NativeCallContext) NativeMethod(i int) INativeMethodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INativeMethodContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INativeMethodContext)
}

func (s *NativeCallContext) AllNativeArray() []INativeArrayContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INativeArrayContext)(nil)).Elem())
	var tst = make([]INativeArrayContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INativeArrayContext)
		}
	}

	return tst
}

func (s *NativeCallContext) NativeArray(i int) INativeArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INativeArrayContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INativeArrayContext)
}

func (s *NativeCallContext) AllPERIOD() []antlr.TerminalNode {
	return s.GetTokens(BeetlParserPERIOD)
}

func (s *NativeCallContext) PERIOD(i int) antlr.TerminalNode {
	return s.GetToken(BeetlParserPERIOD, i)
}

func (s *NativeCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NativeCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NativeCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterNativeCall(s)
	}
}

func (s *NativeCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitNativeCall(s)
	}
}

func (p *BeetlParser) NativeCall() (localctx INativeCallContext) {
	localctx = NewNativeCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, BeetlParserRULE_nativeCall)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(531)
		p.NativeVarRefChain()
	}
	p.SetState(538)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(536)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case BeetlParserLEFT_PAR:
				{
					p.SetState(532)
					p.NativeMethod()
				}

			case BeetlParserLEFT_SQBR:
				{
					p.SetState(533)
					p.NativeArray()
				}

			case BeetlParserPERIOD:
				{
					p.SetState(534)
					p.Match(BeetlParserPERIOD)
				}
				{
					p.SetState(535)
					p.NativeVarRefChain()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(540)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext())
	}

	return localctx
}

// INativeMethodContext is an interface to support dynamic dispatch.
type INativeMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNativeMethodContext differentiates from other interfaces.
	IsNativeMethodContext()
}

type NativeMethodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNativeMethodContext() *NativeMethodContext {
	var p = new(NativeMethodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_nativeMethod
	return p
}

func (*NativeMethodContext) IsNativeMethodContext() {}

func NewNativeMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NativeMethodContext {
	var p = new(NativeMethodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_nativeMethod

	return p
}

func (s *NativeMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *NativeMethodContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(BeetlParserLEFT_PAR, 0)
}

func (s *NativeMethodContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(BeetlParserRIGHT_PAR, 0)
}

func (s *NativeMethodContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *NativeMethodContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NativeMethodContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(BeetlParserCOMMA)
}

func (s *NativeMethodContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(BeetlParserCOMMA, i)
}

func (s *NativeMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NativeMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NativeMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterNativeMethod(s)
	}
}

func (s *NativeMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitNativeMethod(s)
	}
}

func (p *BeetlParser) NativeMethod() (localctx INativeMethodContext) {
	localctx = NewNativeMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, BeetlParserRULE_nativeMethod)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(541)
		p.Match(BeetlParserLEFT_PAR)
	}
	p.SetState(550)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeetlParserLEFT_BRACE)|(1<<BeetlParserLEFT_PAR)|(1<<BeetlParserLEFT_SQBR)|(1<<BeetlParserINCREASE)|(1<<BeetlParserDECREASE)|(1<<BeetlParserADD)|(1<<BeetlParserMIN))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(BeetlParserNOT-40))|(1<<(BeetlParserAT-40))|(1<<(BeetlParserNULL-40))|(1<<(BeetlParserTRUE-40))|(1<<(BeetlParserFALSE-40))|(1<<(BeetlParserDecimalLiteral-40))|(1<<(BeetlParserFloatingPointLiteral-40))|(1<<(BeetlParserStringLiteral-40))|(1<<(BeetlParserIdentifier-40)))) != 0) {
		{
			p.SetState(542)
			p.expression(0)
		}
		p.SetState(547)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == BeetlParserCOMMA {
			{
				p.SetState(543)
				p.Match(BeetlParserCOMMA)
			}
			{
				p.SetState(544)
				p.expression(0)
			}

			p.SetState(549)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(552)
		p.Match(BeetlParserRIGHT_PAR)
	}

	return localctx
}

// INativeArrayContext is an interface to support dynamic dispatch.
type INativeArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNativeArrayContext differentiates from other interfaces.
	IsNativeArrayContext()
}

type NativeArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNativeArrayContext() *NativeArrayContext {
	var p = new(NativeArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_nativeArray
	return p
}

func (*NativeArrayContext) IsNativeArrayContext() {}

func NewNativeArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NativeArrayContext {
	var p = new(NativeArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_nativeArray

	return p
}

func (s *NativeArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *NativeArrayContext) LEFT_SQBR() antlr.TerminalNode {
	return s.GetToken(BeetlParserLEFT_SQBR, 0)
}

func (s *NativeArrayContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NativeArrayContext) RIGHT_SQBR() antlr.TerminalNode {
	return s.GetToken(BeetlParserRIGHT_SQBR, 0)
}

func (s *NativeArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NativeArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NativeArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterNativeArray(s)
	}
}

func (s *NativeArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitNativeArray(s)
	}
}

func (p *BeetlParser) NativeArray() (localctx INativeArrayContext) {
	localctx = NewNativeArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, BeetlParserRULE_nativeArray)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(554)
		p.Match(BeetlParserLEFT_SQBR)
	}
	{
		p.SetState(555)
		p.expression(0)
	}
	{
		p.SetState(556)
		p.Match(BeetlParserRIGHT_SQBR)
	}

	return localctx
}

// INativeVarRefChainContext is an interface to support dynamic dispatch.
type INativeVarRefChainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNativeVarRefChainContext differentiates from other interfaces.
	IsNativeVarRefChainContext()
}

type NativeVarRefChainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNativeVarRefChainContext() *NativeVarRefChainContext {
	var p = new(NativeVarRefChainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_nativeVarRefChain
	return p
}

func (*NativeVarRefChainContext) IsNativeVarRefChainContext() {}

func NewNativeVarRefChainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NativeVarRefChainContext {
	var p = new(NativeVarRefChainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_nativeVarRefChain

	return p
}

func (s *NativeVarRefChainContext) GetParser() antlr.Parser { return s.parser }

func (s *NativeVarRefChainContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(BeetlParserIdentifier)
}

func (s *NativeVarRefChainContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(BeetlParserIdentifier, i)
}

func (s *NativeVarRefChainContext) AllPERIOD() []antlr.TerminalNode {
	return s.GetTokens(BeetlParserPERIOD)
}

func (s *NativeVarRefChainContext) PERIOD(i int) antlr.TerminalNode {
	return s.GetToken(BeetlParserPERIOD, i)
}

func (s *NativeVarRefChainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NativeVarRefChainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NativeVarRefChainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterNativeVarRefChain(s)
	}
}

func (s *NativeVarRefChainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitNativeVarRefChain(s)
	}
}

func (p *BeetlParser) NativeVarRefChain() (localctx INativeVarRefChainContext) {
	localctx = NewNativeVarRefChainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, BeetlParserRULE_nativeVarRefChain)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(558)
		p.Match(BeetlParserIdentifier)
	}
	p.SetState(563)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(559)
				p.Match(BeetlParserPERIOD)
			}
			{
				p.SetState(560)
				p.Match(BeetlParserIdentifier)
			}

		}
		p.SetState(565)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext())
	}

	return localctx
}

// IJsonContext is an interface to support dynamic dispatch.
type IJsonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJsonContext differentiates from other interfaces.
	IsJsonContext()
}

type JsonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonContext() *JsonContext {
	var p = new(JsonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_json
	return p
}

func (*JsonContext) IsJsonContext() {}

func NewJsonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonContext {
	var p = new(JsonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_json

	return p
}

func (s *JsonContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonContext) LEFT_SQBR() antlr.TerminalNode {
	return s.GetToken(BeetlParserLEFT_SQBR, 0)
}

func (s *JsonContext) RIGHT_SQBR() antlr.TerminalNode {
	return s.GetToken(BeetlParserRIGHT_SQBR, 0)
}

func (s *JsonContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *JsonContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *JsonContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(BeetlParserCOMMA)
}

func (s *JsonContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(BeetlParserCOMMA, i)
}

func (s *JsonContext) LEFT_BRACE() antlr.TerminalNode {
	return s.GetToken(BeetlParserLEFT_BRACE, 0)
}

func (s *JsonContext) RIGHT_BRACE() antlr.TerminalNode {
	return s.GetToken(BeetlParserRIGHT_BRACE, 0)
}

func (s *JsonContext) AllJsonKeyValue() []IJsonKeyValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IJsonKeyValueContext)(nil)).Elem())
	var tst = make([]IJsonKeyValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonKeyValueContext)
		}
	}

	return tst
}

func (s *JsonContext) JsonKeyValue(i int) IJsonKeyValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonKeyValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonKeyValueContext)
}

func (s *JsonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JsonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterJson(s)
	}
}

func (s *JsonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitJson(s)
	}
}

func (p *BeetlParser) Json() (localctx IJsonContext) {
	localctx = NewJsonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, BeetlParserRULE_json)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(590)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeetlParserLEFT_SQBR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(566)
			p.Match(BeetlParserLEFT_SQBR)
		}
		p.SetState(575)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeetlParserLEFT_BRACE)|(1<<BeetlParserLEFT_PAR)|(1<<BeetlParserLEFT_SQBR)|(1<<BeetlParserINCREASE)|(1<<BeetlParserDECREASE)|(1<<BeetlParserADD)|(1<<BeetlParserMIN))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(BeetlParserNOT-40))|(1<<(BeetlParserAT-40))|(1<<(BeetlParserNULL-40))|(1<<(BeetlParserTRUE-40))|(1<<(BeetlParserFALSE-40))|(1<<(BeetlParserDecimalLiteral-40))|(1<<(BeetlParserFloatingPointLiteral-40))|(1<<(BeetlParserStringLiteral-40))|(1<<(BeetlParserIdentifier-40)))) != 0) {
			{
				p.SetState(567)
				p.expression(0)
			}
			p.SetState(572)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == BeetlParserCOMMA {
				{
					p.SetState(568)
					p.Match(BeetlParserCOMMA)
				}
				{
					p.SetState(569)
					p.expression(0)
				}

				p.SetState(574)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(577)
			p.Match(BeetlParserRIGHT_SQBR)
		}

	case BeetlParserLEFT_BRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(578)
			p.Match(BeetlParserLEFT_BRACE)
		}
		p.SetState(587)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BeetlParserStringLiteral || _la == BeetlParserIdentifier {
			{
				p.SetState(579)
				p.JsonKeyValue()
			}
			p.SetState(584)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == BeetlParserCOMMA {
				{
					p.SetState(580)
					p.Match(BeetlParserCOMMA)
				}
				{
					p.SetState(581)
					p.JsonKeyValue()
				}

				p.SetState(586)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(589)
			p.Match(BeetlParserRIGHT_BRACE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IJsonKeyValueContext is an interface to support dynamic dispatch.
type IJsonKeyValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJsonKeyValueContext differentiates from other interfaces.
	IsJsonKeyValueContext()
}

type JsonKeyValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonKeyValueContext() *JsonKeyValueContext {
	var p = new(JsonKeyValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_jsonKeyValue
	return p
}

func (*JsonKeyValueContext) IsJsonKeyValueContext() {}

func NewJsonKeyValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonKeyValueContext {
	var p = new(JsonKeyValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_jsonKeyValue

	return p
}

func (s *JsonKeyValueContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonKeyValueContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(BeetlParserStringLiteral, 0)
}

func (s *JsonKeyValueContext) COLON() antlr.TerminalNode {
	return s.GetToken(BeetlParserCOLON, 0)
}

func (s *JsonKeyValueContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *JsonKeyValueContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BeetlParserIdentifier, 0)
}

func (s *JsonKeyValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonKeyValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JsonKeyValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterJsonKeyValue(s)
	}
}

func (s *JsonKeyValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitJsonKeyValue(s)
	}
}

func (p *BeetlParser) JsonKeyValue() (localctx IJsonKeyValueContext) {
	localctx = NewJsonKeyValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, BeetlParserRULE_jsonKeyValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(598)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeetlParserStringLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(592)
			p.Match(BeetlParserStringLiteral)
		}
		{
			p.SetState(593)
			p.Match(BeetlParserCOLON)
		}
		{
			p.SetState(594)
			p.expression(0)
		}

	case BeetlParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(595)
			p.Match(BeetlParserIdentifier)
		}
		{
			p.SetState(596)
			p.Match(BeetlParserCOLON)
		}
		{
			p.SetState(597)
			p.expression(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(BeetlParserDecimalLiteral, 0)
}

func (s *LiteralContext) FloatingPointLiteral() antlr.TerminalNode {
	return s.GetToken(BeetlParserFloatingPointLiteral, 0)
}

func (s *LiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(BeetlParserStringLiteral, 0)
}

func (s *LiteralContext) BooleanLiteral() IBooleanLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *LiteralContext) NULL() antlr.TerminalNode {
	return s.GetToken(BeetlParserNULL, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *BeetlParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, BeetlParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(605)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeetlParserDecimalLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(600)
			p.Match(BeetlParserDecimalLiteral)
		}

	case BeetlParserFloatingPointLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(601)
			p.Match(BeetlParserFloatingPointLiteral)
		}

	case BeetlParserStringLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(602)
			p.Match(BeetlParserStringLiteral)
		}

	case BeetlParserTRUE, BeetlParserFALSE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(603)
			p.BooleanLiteral()
		}

	case BeetlParserNULL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(604)
			p.Match(BeetlParserNULL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(BeetlParserTRUE, 0)
}

func (s *BooleanLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(BeetlParserFALSE, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (p *BeetlParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, BeetlParserRULE_booleanLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(607)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BeetlParserTRUE || _la == BeetlParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeetlParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeetlParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(BeetlParserLEFT_PAR, 0)
}

func (s *ArgumentsContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(BeetlParserRIGHT_PAR, 0)
}

func (s *ArgumentsContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeetlParserListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *BeetlParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, BeetlParserRULE_arguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(609)
		p.Match(BeetlParserLEFT_PAR)
	}
	p.SetState(611)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeetlParserLEFT_BRACE)|(1<<BeetlParserLEFT_PAR)|(1<<BeetlParserLEFT_SQBR)|(1<<BeetlParserINCREASE)|(1<<BeetlParserDECREASE)|(1<<BeetlParserADD)|(1<<BeetlParserMIN))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(BeetlParserNOT-40))|(1<<(BeetlParserAT-40))|(1<<(BeetlParserNULL-40))|(1<<(BeetlParserTRUE-40))|(1<<(BeetlParserFALSE-40))|(1<<(BeetlParserDecimalLiteral-40))|(1<<(BeetlParserFloatingPointLiteral-40))|(1<<(BeetlParserStringLiteral-40))|(1<<(BeetlParserIdentifier-40)))) != 0) {
		{
			p.SetState(610)
			p.ExpressionList()
		}

	}
	{
		p.SetState(613)
		p.Match(BeetlParserRIGHT_PAR)
	}

	return localctx
}

func (p *BeetlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 31:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *BeetlParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
