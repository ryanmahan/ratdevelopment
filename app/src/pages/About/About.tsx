import * as React from "react";
import "../../sass/custom-bulma.scss";
import "./About.css";

export class About extends React.Component<{}, {}> {
    constructor(props: any) {
        super(props);
    }

    render() {
        return <div className="container">
            <h1 className="our-team">Our team</h1>
            <div className="item">
                <div className="info">
                    <h1 className="memberNames">Ryan Mahan</h1>
                    <p className="paragraphs">Meet Ryan, the Scrum Master of the Rat Development team. He has been an integral part of this project. A senior taking CS 529, he has used his software engineering project management skills to lead our team to success. </p>
                    <h1 className="memberNames">Shivani Verma</h1>
                    <p className="paragraphs">Meet Shivani. She's a junior here at UMass. Her hobbies include listening to
                        music and swimming.</p>
                    <h1 className="memberNames">Alexander Marzot</h1>
                    <p className="paragraphs">Meet Alec, a junior at UMass. In his free time, he enjoys coding, rock climbing, and playing board games. </p>
                    <h1 className="memberNames">Artur Tkachenko</h1>
                    <p className="paragraphs">Meet Artur. He enjoys reading, baking, and playing video games.</p>
                    <h1 className="memberNames">Varun Sathiyaseelan</h1>
                    <p className="paragraphs">Meet Varun. His hobbies include writing, cooking, and playing Dungeons & Dragons.</p>
                    <h1 className="memberNames">Gabriel Nathan</h1>
                    <p className="paragraphs">Meet Gabe, a junior at UMass. He always brings a positive atitude to every team meeting and lightens up the atmosphere. His hobbies include programming and making games.</p>
                    <h1 className="memberNames">Temma Eventov</h1>
                    <p className="paragraphs">Meet Temma. When she's not playing rugby or working in IT, she enjoys surfing and rock climbing.</p>
                    <h1 className="memberNames">Liam Brandt</h1>
                    <p className="paragraphs">Meet Liam, who is a great contributor to our team. His hobbies include biking and playing board games.</p>
                    <h1 className="memberNames">Daniel Cline</h1>
                    <p className="paragraphs">Meet Dan, a sophomore at UMass. He enjoys reading, go karting, and drinking boba tea.</p>
                    <h1 className="memberNames">Sarim Ahmed</h1>
                    <p className="paragraphs">Meet Sarim, a junior at UMass. His hobbies include traveling, eating, and sleeping. Fun fact: he's been to half the states in the country!</p>
                    <h1 className="memberNames">Tanner Marsh</h1>
                    <p className="paragraphs">Meet Tanner. He enjoys peanut butter and jelly hotdogs, Hearthstone, and coding.</p>
                </div>
            </div>
        </div>
    }
}

