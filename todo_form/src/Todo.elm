module Todo exposing (User)
import Html exposing (..)
import Html.Attributes exposing (..)

type alias User =
    { name : String
    , email : String
    , password : String
    , loggedIn : Bool
    }

initialModel : User
initialModel =
    { name = ""
    , email = ""
    , password = ""
    , loggedIn = False
    }

view : User -> Html msg
view user =
    div []
        [ h1 [] [ text "To Do Form" ]
        , Html.form []
            [ div []
                [ text "Task Name"
                , input [ id "taskname", type_ "text" ] []
                ]
            , div []
                [ text "Task Description"
                , input [ id "taskdesc", type_ "text" ] []
                ]
            , div []
                [ text "Task Status"
                , input [ id "taskstat", type_ "text" ] []
                ]
            , div []
                [ button [ type_ "submit" ]
                    [ text "Create Task" ]
                ]
            ]
        ]

main : Html msg
main =
    view initialModel
