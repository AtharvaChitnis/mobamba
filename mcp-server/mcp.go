/*                                 
@@@@@@@@@@    @@@@@@@  @@@@@@@   
@@@@@@@@@@@  @@@@@@@@  @@@@@@@@  
@@! @@! @@!  !@@       @@!  @@@  
!@! !@! !@!  !@!       !@!  @!@  
@!! !!@ @!@  !@!       @!@@!@!   
!@!   ! !@!  !!!       !!@!!!    
!!:     !!:  :!!       !!:       
:!:     :!:  :!:       :!:       
:::     ::    ::: :::   ::       
 :      :     :: :: :   :        
*/                                 


										package mcp
															
												import  (
													"context"	//
													"time"	// 	Timeouts and deadlines
													"fmt"	//	Printing
													"log"	//
													"github.com/modelcontextprotocol/go-sdk/mcp"	//Core MCP Package
													"github.com/google/jsonschema-go/jsonschema"	//JSON schema validation and inference
													"encoding/json"
													"os"	//OS operations
													"os/signal"	//Signal handling for gracful shutdown
													"sync"

												)

												/* Transport : [Stdio : Std i/o communication, Server-Sent Events (SSE) : HHTP based event straming 
												, Command Transport : LAunches processes as servers, Custom : WebScoket, gRPC,named, pipes,etc.
												*/
																						
